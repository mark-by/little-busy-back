package microservices

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/scheduler/pkg/proto/scheduler"
	"time"
)

func convertEvent(event *entity.Event) *scheduler.Event {
	if event == nil {
		return nil
	}
	return &scheduler.Event{
		ID:               event.ID,
		CustomerID:       event.CustomerID,
		StartTime:        event.StartTime.Format(time.RFC3339),
		EndTime:          event.EndTime.Format(time.RFC3339),
		Description:      event.Description,
		IsRecurring:      event.IsRecurring,
		Period:           event.Period,
		RecurringEndTime: event.RecurringEndTime.Format(time.RFC3339),
	}
}

func convertProtoEvent(event *scheduler.Event) *entity.Event {
	if event == nil {
		return nil
	}

	startTime, _ := time.Parse(time.RFC3339, event.StartTime)
	endTime, _ := time.Parse(time.RFC3339, event.EndTime)
	recurringEndTime, _ := time.Parse(time.RFC3339, event.RecurringEndTime)
	return &entity.Event{
		ID:               event.ID,
		CustomerID:       event.CustomerID,
		StartTime:        startTime,
		EndTime:          endTime,
		Description:      event.Description,
		IsRecurring:      event.IsRecurring,
		Period:           event.Period,
		RecurringEndTime: recurringEndTime,
	}
}

func convertEvents(events []entity.Event) *scheduler.Events {
	protoEvents := new(scheduler.Events)
	for _, event := range events {
		protoEvents.Events = append(protoEvents.Events, convertEvent(&event))
	}
	return protoEvents
}

func convertProtoEvents(protoEvents *scheduler.Events) []entity.Event {
	if protoEvents == nil {
		return nil
	}

	events := make([]entity.Event, 0, len(protoEvents.Events))
	for _, event := range protoEvents.Events {
		events = append(events, *convertProtoEvent(event))
	}

	return events
}
