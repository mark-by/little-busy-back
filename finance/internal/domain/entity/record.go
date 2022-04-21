package entity

import "time"

type Record struct {
	CustomerID  *int64
	EventID     *int64
	Type        string
	Value       float32
	DateTime    time.Time
	Description *string
}
