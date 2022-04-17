package microservices

import (
	"context"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/mark-by/little-busy-back/finance/pkg/proto/finance"
	"github.com/mark-by/little-busy-back/pkg/converter"
	"time"
)

type Record struct {
	financeClient finance.FinanceServiceClient
}

func (r Record) Create(record *entity.Record) error {
	financeRecord, err := converter.ConvertStruct[entity.Record, finance.Record](*record)
	if err != nil {
		return err
	}
	_, err = r.financeClient.Create(context.Background(), &financeRecord)
	return err
}

func (r Record) GetRecordsForDay(date time.Time) ([]entity.Record, error) {
	records, err := r.financeClient.GetRecordsForDay(context.Background(), &finance.DateRequest{
		Year:  int32(date.Year()),
		Month: int32(date.Month()),
		Day:   int32(date.Day()),
	})
	if err != nil {
		return nil, err
	}

	newRecords, err := converter.ConvertPointerSlice[finance.Record, entity.Record](records.Records)
	if err != nil {
		return nil, err
	}

	return newRecords, nil
}

func NewRecord(financeClient finance.FinanceServiceClient) *Record {
	return &Record{financeClient: financeClient}
}

var _ repository.Record = &Record{}
