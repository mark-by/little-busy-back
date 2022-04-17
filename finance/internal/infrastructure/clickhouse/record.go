package clickhouse

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/mark-by/little-busy-back/finance/internal/domain/entity"
	"github.com/mark-by/little-busy-back/finance/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type Record struct {
	db clickhouse.Conn
}

func (r Record) Create(record *entity.Record) error {
	err := r.db.Exec(context.Background(),
		"insert into Records (ClientID, EventID, Type, Value, DateTime, Description) "+
			"values ($1, $2, $3, $4, $5, $6)",
		record.ClientID,
		record.EventID,
		record.Type,
		record.Value,
		record.DateTime,
		record.Description)
	if err != nil {
		return errors.Wrap(err, "fail to insert record")
	}
	return nil
}

func (r Record) GetRecordsForDay(date time.Time) ([]entity.Record, error) {
	var records []entity.Record
	err := r.db.Select(context.Background(), &records,
		"select * from Records where toDate(DateTime) == toDate($1)", date)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get records for the day")
	}
	panic("implement me")
}

func NewRecord(db clickhouse.Conn) *Record {
	return &Record{db}
}

var _ repository.Record = &Record{}
