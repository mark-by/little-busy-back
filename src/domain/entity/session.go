package entity

import "time"

type Session struct {
	ID         int64
	Customer   *Customer
	Datetime   time.Time
	Duration   time.Duration
	IsConstant bool
	IsActive   bool
}
