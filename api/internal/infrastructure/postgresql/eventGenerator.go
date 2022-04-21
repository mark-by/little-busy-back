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

func (e eventGenerator) timeCheck(recurringEvent entity.Event, curr time.Time) bool {
	return !dateIsLess(curr, recurringEvent.StartTime) && (recurringEvent.RecurringEndTime == nil || !dateIsLess(*recurringEvent.RecurringEndTime, curr))
}

func dateIsLess(l, r time.Time) bool {
	ly, lm, ld := l.Date()
	ry, rm, rd := r.Date()
	if ly > ry {
		return false
	}
	if lm > rm {
		return false
	}
	if ld >= rd {
		return false
	}
	return true
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
