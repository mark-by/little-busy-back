package entity

import "time"

type Event struct {
	ClientTel string
	StartTime time.Time
	EndTime   time.Time
	Price     float64
}
