package microservices

import (
	"context"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	protoScheduler "github.com/mark-by/little-busy-back/scheduler/pkg/proto/scheduler"
	"time"
)

type Events struct {
	schedulerClient protoScheduler.SchedulerClient
}

func (e Events) DeleteAllForCustomer(customerID int64) error {
	_, err := e.schedulerClient.DeleteAllForCustomer(context.Background(), &protoScheduler.CustomerID{ID: customerID})
	return err
}

func (e Events) Get(eventID int64) (*entity.Event, error) {
	event, err := e.schedulerClient.GetEvent(context.Background(), &protoScheduler.EventID{EventID: eventID})
	return convertProtoEvent(event), err
}

func (e Events) GetForMonth(year, month int) ([]entity.Event, error) {
	events, err := e.schedulerClient.GetEventsFor(context.Background(), &protoScheduler.Date{
		Year:  int32(year),
		Month: int32(month),
	})

	return convertProtoEvents(events), err
}

func (e Events) GetForDay(year, month, day int) ([]entity.Event, error) {
	events, err := e.schedulerClient.GetEventsFor(context.Background(), &protoScheduler.Date{
		Year:  int32(year),
		Month: int32(month),
		Day:   int32(day),
	})

	return convertProtoEvents(events), err
}

func (e Events) GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error) {
	events, err := e.schedulerClient.GetEventsForCustomer(context.Background(), &protoScheduler.CustomerRequest{
		CustomerID: customerID,
		Since:      since.Unix(),
		Days:       int32(days),
	})

	return convertProtoEvents(events), err
}

func (e Events) Create(event *entity.Event) (*entity.Event, error) {
	newEvent, err := e.schedulerClient.CreateEvent(context.Background(), convertEvent(event))
	return convertProtoEvent(newEvent), err
}

func (e Events) Update(event *entity.Event, currStartTime time.Time, withNext bool) error {
	_, err := e.schedulerClient.UpdateEvent(context.Background(), &protoScheduler.UpdateRequest{
		Date:     currStartTime.Unix(),
		WithNext: withNext,
		Event:    convertEvent(event),
	})
	return err
}

func (e Events) Delete(eventID int64, currStartTime time.Time, withNext bool) error {
	_, err := e.schedulerClient.DeleteEvent(context.Background(), &protoScheduler.DeleteRequest{
		Date:     currStartTime.Unix(),
		WithNext: withNext,
		EventID:  eventID,
	})
	return err
}

func NewEvents(client protoScheduler.SchedulerClient) *Events {
	return &Events{schedulerClient: client}
}

var _ repository.Events = &Events{}
