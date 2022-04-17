package grpc

import (
	"context"
	"github.com/mark-by/little-busy-back/scheduler/internal/application"
	"github.com/mark-by/little-busy-back/scheduler/internal/domain/entity"
	protoScheduler "github.com/mark-by/little-busy-back/scheduler/pkg/proto/scheduler"
	"go.uber.org/zap"
	"time"
)

type SchedulerService struct {
	eventsApp application.EventsI
	logger    *zap.SugaredLogger
}

func (s SchedulerService) DeleteAllForCustomer(ctx context.Context, id *protoScheduler.CustomerID) (*protoScheduler.Empty, error) {
	return &protoScheduler.Empty{}, s.eventsApp.DeleteForAll(id.GetID())
}

func NewSchedulerService(eventsApp application.EventsI, logger *zap.SugaredLogger) *SchedulerService {
	return &SchedulerService{
		eventsApp: eventsApp,
		logger:    logger,
	}
}

func (s SchedulerService) GetEvent(ctx context.Context, id *protoScheduler.EventID) (*protoScheduler.Event, error) {
	event, err := s.eventsApp.Get(id.EventID)
	if err != nil {
		s.logger.With("id", id.EventID).Error("fail to get event: ", err)
		return nil, err
	}
	return convertEvent(event), nil
}

func (s SchedulerService) GetEventsFor(ctx context.Context, date *protoScheduler.Date) (*protoScheduler.Events, error) {
	var events entity.Events
	var err error
	if date.Day != 0 {
		events, err = s.eventsApp.GetForDay(int(date.Year), int(date.Month), int(date.Day))
	} else {
		events, err = s.eventsApp.GetForMonth(int(date.Year), int(date.Month))
	}

	if err != nil {
		s.logger.With("date", date).Error("fail to get events for date: ", err)
		return nil, err
	}

	return convertEvents(events), nil
}

func (s SchedulerService) GetEventsForCustomer(ctx context.Context, request *protoScheduler.CustomerRequest) (*protoScheduler.Events, error) {
	events, err := s.eventsApp.GetForCustomer(request.CustomerID, time.Unix(request.Since, 0), int(request.Days))
	if err != nil {
		s.logger.With("customer", request.CustomerID,
			"since", time.Unix(request.Since, 0),
			"days", request.Days).Error("fail to get events for customer: ", err)
		return nil, err
	}

	return convertEvents(events), nil
}

func (s SchedulerService) CreateEvent(ctx context.Context, event *protoScheduler.Event) (*protoScheduler.Event, error) {
	newEvent, err := s.eventsApp.Create(convertProtoEvent(event))
	if err != nil {
		s.logger.With("event", event).Error("fail to create event: ", err)
		return nil, err
	}
	return convertEvent(newEvent), nil
}

func (s SchedulerService) UpdateEvent(ctx context.Context, request *protoScheduler.UpdateRequest) (*protoScheduler.Empty, error) {
	err := s.eventsApp.Update(convertProtoEvent(request.Event), time.Unix(request.Date, 0), request.WithNext)
	if err != nil {
		s.logger.With("event", request.Event).Error("fail to update event: ", err)
		return new(protoScheduler.Empty), err
	}
	return new(protoScheduler.Empty), nil
}

func (s SchedulerService) DeleteEvent(ctx context.Context, request *protoScheduler.DeleteRequest) (*protoScheduler.Empty, error) {
	err := s.eventsApp.Delete(request.EventID, time.Unix(request.Date, 0), request.WithNext)
	if err != nil {
		s.logger.With("event", request.EventID).Error("fail to delete event: ", err)
		return new(protoScheduler.Empty), err
	}

	return new(protoScheduler.Empty), nil
}

var _ protoScheduler.SchedulerServer = &SchedulerService{}
