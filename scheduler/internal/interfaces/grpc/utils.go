package grpc

import (
	"github.com/mark-by/little-busy-back/pkg/utils"
	"github.com/mark-by/little-busy-back/scheduler/internal/domain/entity"
	"github.com/mark-by/little-busy-back/scheduler/pkg/proto/scheduler"
	"time"
)

func convertEvent(event *entity.Event) *scheduler.Event {
	if event == nil {
		return nil
	}
	var recurringEndTime string
	if event.RecurringEndTime != nil {
		recurringEndTime = event.RecurringEndTime.Format(time.RFC3339)
	}
	return &scheduler.Event{
		ID:               event.ID,
		CustomerID:       utils.DropNil(event.CustomerID),
		StartTime:        event.StartTime.Format(time.RFC3339),
		EndTime:          event.EndTime.Format(time.RFC3339),
		Description:      utils.DropNil(event.Description),
		IsRecurring:      event.IsRecurring,
		Period:           utils.DropNil(event.Period),
		RecurringEndTime: recurringEndTime,
	}
}

func convertProtoEvent(event *scheduler.Event) *entity.Event {
	if event == nil {
		return nil
	}
	var recurringEndTime *time.Time
	if event.RecurringEndTime != "" {
		recurringEndTimeTmp, _ := time.Parse(time.RFC3339, event.RecurringEndTime)
		recurringEndTime = &recurringEndTimeTmp
	}

	startTime, _ := time.Parse(time.RFC3339, event.StartTime)
	endTime, _ := time.Parse(time.RFC3339, event.EndTime)
	return &entity.Event{
		ID:               event.ID,
		CustomerID:       &event.CustomerID,
		StartTime:        startTime,
		EndTime:          endTime,
		Description:      &event.Description,
		IsRecurring:      event.IsRecurring,
		Period:           &event.Period,
		RecurringEndTime: recurringEndTime,
	}
}

func convertEvents(events entity.Events) *scheduler.Events {
	protoEvents := new(scheduler.Events)
	for _, event := range events {
		protoEvents.Events = append(protoEvents.Events, convertEvent(&event))
	}
	return protoEvents
}
