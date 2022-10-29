package repository

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"time"
)

type Record interface {
	Create(record *entity.Record) error
	Delete(recordID int64) error
	GetRecordsForDay(date time.Time) ([]entity.Record, error)
	GetStatForMonth(year, month int) ([]entity.Value, error)
	GetStatForYear(year int) ([]entity.Value, error)
	GetProfit(start, end time.Time) (float32, error)
	SaveBatch(records []entity.Record) error
	Select(since int64, limit int) ([]entity.Record, error)
	Update(recordID int64, value float32) error
}
