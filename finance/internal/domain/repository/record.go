package repository

import (
	"github.com/mark-by/little-busy-back/finance/internal/domain/entity"
	"time"
)

type Record interface {
	Create(record *entity.Record) error
	GetRecordsForDay(date time.Time) ([]entity.Record, error)
}
