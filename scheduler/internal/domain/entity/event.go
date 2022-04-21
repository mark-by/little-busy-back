package entity

import (
	"sort"
	"time"
)

type Event struct {
	ID               int64      `json:"id" db:"id"`
	StartTime        time.Time  `json:"start_time" validate:"required" db:"start_time"`
	EndTime          time.Time  `json:"end_time" validate:"gtfield=StartTime" db:"end_time"`
	CustomerID       *int64     `json:"customer_id,omitempty" validate:"gt=0" db:"customer_id"`
	Description      *string    `json:"description" db:"description"`
	IsRecurring      bool       `json:"is_recurring" db:"is_recurring"`
	Period           *string    `json:"period,omitempty" validate:"oneof=daily weekly monthly ''" db:"period"`
	RecurringEndTime *time.Time `json:"recurring_end_time,omitempty" db:"recurring_end_time"`
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
