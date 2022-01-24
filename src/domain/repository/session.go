package repository

import (
	"github.com/mark-by/little-busy-back/src/domain/entity"
	"time"
)

type Session interface {
	Get(ID int64) (*entity.Session, error)
	Save(session *entity.Session) error
	Delete(ID int64) error
	GetForMonth(date time.Time) ([]entity.Session, error)
	GetForCustomer(customerID int64) ([]entity.Session, error)
	FindFreeSpace(duration time.Duration, start, end time.Time) ([]entity.Session, error)
}
