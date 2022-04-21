package postgresql

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type Record struct {
	db *pgxpool.Pool
}

func (r Record) Create(record *entity.Record) error {
	_, err := r.db.Exec(context.Background(),
		`insert into records (customer_id, event_id, is_income, value, datetime, description)
		values ($1, $2, $3, $4, $5, $6)`,
		record.CustomerID,
		record.EventID,
		record.IsIncome,
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
	err := pgxscan.Select(context.Background(), r.db, &records,
		"select * from records where datetime::date == $1::date", date)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get records for the day")
	}
	panic("implement me")
}

func (v Value) GetStatForMonth(year, month int) ([]entity.Value, error) {
	rows, err := v.db.Query(context.Background(),
		`SELECT
					toDate(DateTime) as date,
    				sum(if(Type = 'income', Value, 0)) AS incomes,
    				sum(if(Type = 'cost', Value, 0)) AS costs
				FROM Records
				GROUP BY toDate(DateTime)
				HAVING toYear(DateTime) = $1 and toMonth(DateTime) = $2`, year, month)

	if err != nil {
		return nil, errors.Wrap(err, "fail to select for month")
	}
	defer rows.Close()

	var values []entity.Value
	for rows.Next() {
		var value entity.Value

		err = rows.Scan(&value.Date, &value.Incomes, &value.Costs)
		if err != nil {
			return nil, errors.Wrap(err, "fail to scan row in select for month")
		}

		values = append(values, value)
	}

	return values, nil
}

func (v Value) GetStatForYear(year int) ([]entity.Value, error) {
	rows, err := v.db.Query(context.Background(),
		`SELECT
					to_char(datetime, 'mm')::int as month,
    				sum(case when is_income then value else 0 end) AS incomes,
    				sum(case when not is_income then value else 0 end) AS costs
				FROM records
				GROUP BY to_char(datetime, 'mm')::int
				HAVING to_char(datetime, 'yyyy')::int = $1`, year)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select for year")
	}
	defer rows.Close()

	var values []entity.Value
	for rows.Next() {
		var (
			month int
			value entity.Value
		)

		err = rows.Scan(&month, &value.Incomes, &value.Costs)
		if err != nil {
			return nil, errors.Wrap(err, "fail to scan row in select for year")
		}

		value.DateTime = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		values = append(values, value)
	}

	return values, nil
}

func NewRecord(db *pgxpool.Pool) *Record {
	return &Record{db}
}

var _ repository.Record = &Record{}
