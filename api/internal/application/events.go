package application

import (
	"fmt"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type EventsI interface {
	Get(eventID int64) (*entity.Event, error)
	GetForMonth(year, month int) ([]entity.Event, error)
	GetForDay(year, month, day int) ([]entity.Event, error)
	GetNotPaidForDay(year, month, day int) ([]entity.Event, error)
	GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error)
	Create(event *entity.Event) (*entity.Event, error)
	Update(event *entity.Event, currStartTime time.Time, withNext bool) error
	Delete(eventID int64, currStartTime time.Time, withNext bool) error
	DeleteAllForCustomer(customerID int64) error
}

type Events struct {
	repoEvents    repository.Events
	repoCustomers repository.Customers
	repoRecords   repository.Record
}

func (e Events) DeleteAllForCustomer(customerID int64) error {
	return e.repoEvents.DeleteAllForCustomer(customerID)
}

func (e Events) Get(eventID int64) (*entity.Event, error) {
	return e.repoEvents.Get(eventID)
}

func (e Events) GetForMonth(year, month int) ([]entity.Event, error) {
	return e.repoEvents.GetForMonth(year, month)
}

func (e Events) GetForDay(year, month, day int) ([]entity.Event, error) {
	return e.repoEvents.GetForDay(year, month, day)
}

func (e Events) GetNotPaidForDay(year, month, day int) ([]entity.Event, error) {
	events, err := e.GetForDay(year, month, day)
	if err != nil {
		return nil, err
	}
	records, err := e.repoRecords.GetRecordsForDay(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		return nil, err
	}

	paid := map[int64]bool{}
	for _, record := range records {
		if record.EventID == nil {
			continue
		}
		paid[*record.EventID] = true
	}

	var notPaidEvents []entity.Event

	for _, event := range events {
		if _, ok := paid[event.ID]; !ok {
			notPaidEvents = append(notPaidEvents, event)
		}
	}

	return notPaidEvents, nil
}

func (e Events) GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error) {
	return e.repoEvents.GetForCustomer(customerID, since, days)
}

func (e Events) Create(event *entity.Event) (*entity.Event, error) {
	return e.repoEvents.Create(event)
}

func (e Events) Update(event *entity.Event, currStartTime time.Time, withNext bool) error {
	oldEvent, err := e.repoEvents.Get(event.ID)
	if err != nil {
		return err
	}

	if !oldEvent.IsRecurring {
		// не был повторяющимся
		return e.repoEvents.UpdateRegular(event)
	}

	if oldEvent.IsRecurring && !event.IsRecurring {
		// перестал быть повторяющимся
		err = e.repoEvents.DeleteWithNextRecurring(event, currStartTime)
		if err != nil {
			return err
		}
		_, err = e.repoEvents.Create(event)
		return err
	}

	if oldEvent.IsRecurring && event.IsRecurring {
		// повторяющееся
		if withNext {
			if oldEvent.StartTime.Sub(currStartTime) == 0 {
				// the same event
				event.IsRecurring = false
				return e.repoEvents.UpdateRegular(event)
			}
			return e.repoEvents.UpdateWithNextRecurring(event, currStartTime)
		}

		return e.repoEvents.UpdateOnlyCurrRecurring(event, currStartTime)
	}

	return nil
}

func (e Events) Delete(eventID int64, currStartTime time.Time, withNext bool) error {
	event, err := e.repoEvents.Get(eventID)
	if err != nil {
		return err
	}
	if event == nil {
		return errors.New(fmt.Sprintf("no such event with id %v", eventID))
	}

	if event.IsRecurring {
		// recurring
		if withNext {
			return e.repoEvents.DeleteWithNextRecurring(event, currStartTime)
		}
		if event.StartTime.Sub(currStartTime) == 0 {
			// the same event
			// передвигаем начало повторяющегося событие
			nextEvent := event.NextRecurring()
			if nextEvent == nil {
				return e.repoEvents.DeleteOnlyCurrRecurring(event, currStartTime)
			}
			nextEvent.IsRecurring = false // костыль
			return e.repoEvents.UpdateRegular(nextEvent)
		}
		return e.repoEvents.DeleteOnlyCurrRecurring(event, currStartTime)
	}

	return e.repoEvents.DeleteRegular(eventID)
}

func NewEvents(events repository.Events, customers repository.Customers, records repository.Record) *Events {
	return &Events{repoEvents: events, repoCustomers: customers, repoRecords: records}
}

var _ EventsI = &Events{}
