package repository

import (
	"github.com/mark-by/little-busy-back/scheduler/internal/domain/entity"
	"time"
)

type Event interface {
	Get(eventID int64) (*entity.Event, error)
	GetForMonth(year, month int) ([]entity.Event, error)
	GetForDay(year, month, day int) ([]entity.Event, error)
	GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error)
	Create(event *entity.Event) (*entity.Event, error)
	UpdateRegular(event *entity.Event) error
	UpdateWithNextRecurring(event *entity.Event, currStartTime time.Time) error
	UpdateOnlyCurrRecurring(event *entity.Event, currStartTime time.Time) error
	DeleteRegular(eventID int64) error
	DeleteAllForCustomer(customerID int64) error
	DeleteWithNextRecurring(event *entity.Event, currStartTime time.Time) error
	DeleteOnlyCurrRecurring(event *entity.Event, currStartTime time.Time) error
}
