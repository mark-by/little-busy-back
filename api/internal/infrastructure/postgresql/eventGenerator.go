package postgresql

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"time"
)

type eventGenerator struct {
	generatedEvents entity.Events
	eventPicker     *eventPicker
}

func newEventGenerator(picker *eventPicker) *eventGenerator {
	return &eventGenerator{
		generatedEvents: nil,
		eventPicker:     picker,
	}
}

func (e *eventGenerator) addEventToStoreForTime(recurringEvent entity.Event, curr time.Time) {
	if !e.timeCheck(recurringEvent, curr) {
		return
	}
	newEvent := recurringEvent.CopyWithNewDate(curr)
	newEvent.IsRecurring = true
	e.generatedEvents = append(e.generatedEvents, newEvent)
}

// check dates without time
func (e eventGenerator) timeCheck(recurringEvent entity.Event, curr time.Time) bool {
	curr = time.Date(curr.Year(), curr.Month(), curr.Day(), 0, 0, 0, 0, time.Local)
	recurringEventStartTime := time.Date(recurringEvent.StartTime.Year(), recurringEvent.StartTime.Month(), recurringEvent.StartTime.Day(),
		0, 0, 0, 0, time.Local)
	var recurringEventEndTime time.Time
	if recurringEvent.RecurringEndTime != nil {
		recurringEventEndTime = time.Date(recurringEvent.RecurringEndTime.Year(), recurringEvent.RecurringEndTime.Month(), recurringEvent.RecurringEndTime.Day(),
			0, 0, 0, 0, time.Local)
	}
	return !dateIsLess(curr, recurringEventStartTime) && (recurringEvent.RecurringEndTime == nil || !dateIsLess(recurringEventEndTime, curr))
}

// l < r
func dateIsLess(l, r time.Time) bool {
	return l.Sub(r).Hours() < 0
}

func (e eventGenerator) GenEvents(from, to time.Time) entity.Events {
	curr := from
	for to.Sub(curr).Hours() > 0 {
		for _, event := range e.eventPicker.GetDaily() {
			e.addEventToStoreForTime(event, curr)
		}
		for _, event := range e.eventPicker.GetForWeek(int(curr.Weekday())) {
			e.addEventToStoreForTime(event, curr)
		}
		for _, event := range e.eventPicker.GetForMontDay(curr.Day()) {
			e.addEventToStoreForTime(event, curr)
		}

		curr = curr.AddDate(0, 0, 1)
	}

	return e.generatedEvents
}
