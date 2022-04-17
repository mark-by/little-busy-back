package application

import (
	"github.com/mark-by/little-busy-back/finance/internal/domain/entity"
	"github.com/mark-by/little-busy-back/finance/internal/domain/repository"
	"time"
)

type RecordI interface {
	Create(record *entity.Record) error
	GetRecordsForDay(date time.Time) ([]entity.Record, error)
}

type Record struct {
	records repository.Record
}

func (r Record) Create(record *entity.Record) error {
	return r.records.Create(record)
}

func (r Record) GetRecordsForDay(date time.Time) ([]entity.Record, error) {
	return r.records.GetRecordsForDay(date)
}

func NewRecord(records repository.Record) *Record {
	return &Record{records}
}

var _ RecordI = &Record{}
