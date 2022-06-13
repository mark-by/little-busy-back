package entity

import "time"

type Record struct {
	ID          int64     `json:"id" db:"id"`
	CustomerID  *int64    `json:"customer_id" db:"customer_id"`
	EventID     *int64    `json:"event_id" db:"event_id"`
	IsIncome    bool      `json:"is_income" db:"is_income"`
	Value       float32   `json:"value" db:"value"`
	DateTime    time.Time `json:"datetime" db:"datetime"`
	Description *string   `json:"description" db:"description"`
}
