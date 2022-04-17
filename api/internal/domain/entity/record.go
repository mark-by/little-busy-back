package entity

import "time"

type Record struct {
	ClientID    int64
	EventID     int64
	Type        string
	Value       float32
	DateTime    time.Time
	Description string
}
