package entity

import "time"

type Record struct {
	CustomerID  *int64
	EventID     *int64
	IsIncome    bool
	Value       float32
	DateTime    time.Time
	Description *string
}
