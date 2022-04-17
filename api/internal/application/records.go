package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"time"
)

type RecordsI interface {
	Create(record *entity.Record) error
	GetRecordsForDay(date time.Time) ([]entity.Record, error)
}

type Records struct {
	records repository.Record
}

func (r Records) Create(record *entity.Record) error {
	return r.records.Create(record)
}

func (r Records) GetRecordsForDay(date time.Time) ([]entity.Record, error) {
	return r.records.GetRecordsForDay(date)
}

var _ RecordsI = &Records{}
