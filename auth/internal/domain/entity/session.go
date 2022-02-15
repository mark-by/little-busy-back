package entity

import "time"

type Session struct {
	ID         string
	UserID     int
	Expiration time.Time
}
