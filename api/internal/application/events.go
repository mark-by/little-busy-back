package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
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
	event, err := e.repoEvents.Get(eventID)
	if err != nil {
		return nil, err
	}

	customers, err := e.repoCustomers.GetCustomers([]int64{event.CustomerID})
	if err != nil {
		return nil, err
	}

	if len(customers) > 0 {
		event.Customer = &customers[0]
	}
	return event, nil
}

func (e Events) GetForMonth(year, month int) ([]entity.Event, error) {
	events, err := e.repoEvents.GetForMonth(year, month)
	if err != nil {
		return nil, err
	}
	return e.addCustomers(events)
}

func (e Events) GetForDay(year, month, day int) ([]entity.Event, error) {
	events, err := e.repoEvents.GetForDay(year, month, day)
	if err != nil {
		return nil, err
	}
	return e.addCustomers(events)
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

	has := map[int64]bool{}
	for _, record := range records {
		has[record.EventID] = true
	}

	var notPaidEvents []entity.Event

	for _, event := range events {
		if _, ok := has[event.ID]; ok {
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
	return e.repoEvents.Update(event, currStartTime, withNext)
}

func (e Events) Delete(eventID int64, currStartTime time.Time, withNext bool) error {
	return e.repoEvents.Delete(eventID, currStartTime, withNext)
}

func NewEvents(events repository.Events, customers repository.Customers) *Events {
	return &Events{repoEvents: events, repoCustomers: customers}
}

var _ EventsI = &Events{}
