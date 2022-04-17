package application

import (
	"fmt"
	"github.com/mark-by/little-busy-back/scheduler/internal/domain/entity"
	"github.com/mark-by/little-busy-back/scheduler/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type EventsI interface {
	Get(eventID int64) (*entity.Event, error)
	GetForMonth(year, month int) ([]entity.Event, error)
	GetForDay(year, month, day int) ([]entity.Event, error)
	GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error)
	Create(event *entity.Event) (*entity.Event, error)

	Update(event *entity.Event, currStartTime time.Time, withNext bool) error
	Delete(eventID int64, currStartTime time.Time, withNext bool) error
	DeleteForAll(customerID int64) error
}

type Events struct {
	repo repository.Event
}

func (e Events) DeleteForAll(customerID int64) error {
	return e.repo.DeleteAllForCustomer(customerID)
}

func (e Events) Get(eventID int64) (*entity.Event, error) {
	return e.repo.Get(eventID)
}

func (e Events) GetForMonth(year, month int) ([]entity.Event, error) {
	return e.repo.GetForMonth(year, month)
}

func (e Events) GetForDay(year, month, day int) ([]entity.Event, error) {
	return e.repo.GetForDay(year, month, day)
}

func (e Events) GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error) {
	return e.repo.GetForCustomer(customerID, since, days)
}

func (e Events) Create(event *entity.Event) (*entity.Event, error) {
	return e.repo.Create(event)
}

func (e Events) Update(event *entity.Event, currStartTime time.Time, withNext bool) error {
	oldEvent, err := e.repo.Get(event.ID)
	if err != nil {
		return err
	}

	if !oldEvent.IsRecurring {
		// не был повторяющимся
		return e.repo.UpdateRegular(event)
	}

	if oldEvent.IsRecurring && !event.IsRecurring {
		// перестал быть повторяющимся
		err = e.repo.DeleteWithNextRecurring(event, currStartTime)
		if err != nil {
			return err
		}
		_, err = e.repo.Create(event)
		return err
	}

	if oldEvent.IsRecurring && event.IsRecurring {
		// повторяющееся
		if withNext {
			if oldEvent.StartTime.Sub(currStartTime) == 0 {
				// the same event
				event.IsRecurring = false
				return e.repo.UpdateRegular(event)
			}
			return e.repo.UpdateWithNextRecurring(event, currStartTime)
		}

		return e.repo.UpdateOnlyCurrRecurring(event, currStartTime)
	}

	return nil
}

func (e Events) Delete(eventID int64, currStartTime time.Time, withNext bool) error {
	event, err := e.repo.Get(eventID)
	if err != nil {
		return err
	}
	if event == nil {
		return errors.New(fmt.Sprintf("no such event with id %v", eventID))
	}

	if event.IsRecurring {
		// recurring
		if withNext {
			return e.repo.DeleteWithNextRecurring(event, currStartTime)
		}
		if event.StartTime.Sub(currStartTime) == 0 {
			// the same event
			// передвигаем начало повторяющегося событие
			nextEvent := event.NextRecurring()
			if nextEvent == nil {
				return e.repo.DeleteOnlyCurrRecurring(event, currStartTime)
			}
			nextEvent.IsRecurring = false // костыль
			return e.repo.UpdateRegular(nextEvent)
		}
		return e.repo.DeleteOnlyCurrRecurring(event, currStartTime)
	}

	return e.repo.DeleteRegular(eventID)
}

func NewEvents(repo repository.Event) *Events {
	return &Events{repo}
}

var _ EventsI = &Events{}
