package repository

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"time"
)

type Events interface {
	Get(eventID int64) (*entity.Event, error)
	GetForMonth(year, month int) ([]entity.Event, error)
	GetForDay(year, month, day int) ([]entity.Event, error)
	GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error)
	Create(event *entity.Event) (*entity.Event, error)
	Update(event *entity.Event, currStartTime time.Time, withNext bool) error
	Delete(eventID int64, currStartTime time.Time, withNext bool) error
	DeleteAllForCustomer(customerID int64) error
}
