package postgresql

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"sort"
	"time"
)

type recurringEventsAdder struct {
	recurringEventPicker *eventPicker
}

func newRecurringEventsAdder(recurringEvents entity.Events) *recurringEventsAdder {
	picker := newEventPicker()

	for _, event := range recurringEvents {
		if event.Period == nil {
			continue
		}
		switch *event.Period {
		case "daily":
			picker.AddToDaily(event)
		case "weekly":
			picker.AddToWeekDay(int(event.StartTime.Weekday()), event)
		case "monthly":
			picker.AddToMonthDay(event.StartTime.Day(), event)
		}
	}

	return &recurringEventsAdder{
		recurringEventPicker: picker,
	}
}

func (r recurringEventsAdder) AddTo(events entity.Events, from, to time.Time) entity.Events {
	recurringEvents := newEventGenerator(r.recurringEventPicker).GenEvents(from, to)

	events = append(events, recurringEvents...)
	sort.Sort(events)
	return events
}
