package entity

import (
	"sort"
	"time"
)

type Event struct {
	ID               int64      `db:"id"`
	StartTime        time.Time  `db:"start_time"`
	EndTime          time.Time  `db:"end_time"`
	CustomerID       *int64     `db:"customer_id"`
	Description      *string    `db:"description"`
	IsRecurring      bool       `db:"is_recurring"`
	Period           *string    `db:"period"`
	RecurringEndTime *time.Time `db:"recurring_end_time"`
}

func (e Event) CopyWithNewDate(date time.Time) Event {
	newEvent := e
	newYear, newMonth, newDay := date.Date()
	currYear, currMonth, currDay := newEvent.StartTime.Date()
	newEvent.StartTime = newEvent.StartTime.AddDate(newYear-currYear, int(newMonth)-int(currMonth), newDay-currDay)
	newEvent.EndTime = newEvent.EndTime.AddDate(newYear-currYear, int(newMonth)-int(currMonth), newDay-currDay)

	return newEvent
}

func (e Event) NextRecurring() *Event {
	if !e.IsRecurring || e.Period == nil {
		return nil
	}

	var event Event
	switch *e.Period {
	case "daily":
		event = e.CopyWithNewDate(e.StartTime.AddDate(0, 0, 1))
	case "weekly":
		event = e.CopyWithNewDate(e.StartTime.AddDate(0, 0, 7))
	case "monthly":
		event = e.CopyWithNewDate(e.StartTime.AddDate(0, 1, 0))
	}

	if event.recurringIsNotEnded() {
		return &event
	}

	return nil
}

func (e Event) recurringIsNotEnded() bool {
	return e.RecurringEndTime == nil || e.RecurringEndTime.Sub(e.StartTime) >= 0
}

type Events []Event

func (e Events) Len() int {
	return len(e)
}

func (e Events) Less(i, j int) bool {
	return e[i].StartTime.Sub(e[j].StartTime) < 0
}

func (e Events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

var _ sort.Interface = &Events{}
