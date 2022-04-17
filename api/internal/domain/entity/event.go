package entity

import "time"

type Event struct {
	ID               int64     `json:"id"`
	StartTime        time.Time `json:"start_time" validate:"required"`
	EndTime          time.Time `json:"end_time" validate:"required,gtfield=StartTime"`
	CustomerID       int64     `json:"customer_id,omitempty" validate:"gte=0"`
	Customer         *Customer `json:"customer,omitempty"`
	Description      string    `json:"description"`
	IsRecurring      bool      `json:"is_recurring"`
	Period           string    `json:"period,omitempty" validate:"oneof=daily weekly monthly ''"`
	RecurringEndTime time.Time `json:"recurring_end_time,omitempty"`
}
