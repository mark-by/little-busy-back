package postgresql

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type Record struct {
	db *pgxpool.Pool
}

func (r Record) Select(since int64, limit int) ([]entity.Record, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 10 {
		limit = 10
	}

	var records []entity.Record
	var err error
	err = pgxscan.Select(context.Background(), r.db, &records,
		`SELECT * from records
		ORDER BY datetime DESC, id DESC
		OFFSET $1
		LIMIT $2`, since, limit)

	if err != nil {
		return nil, errors.Wrap(err, "fail select from records")
	}

	return records, nil
}

func (r Record) Create(record *entity.Record) error {
	err := r.db.QueryRow(context.Background(),
		`insert into records (customer_id, event_id, is_income, value, datetime, description)
		values ($1, $2, $3, $4, $5, $6) returning id`,
		record.CustomerID,
		record.EventID,
		record.IsIncome,
		record.Value,
		record.DateTime,
		record.Description).Scan(&record.ID)
	if err != nil {
		return errors.Wrap(err, "fail to insert record")
	}
	return nil
}

func (r Record) Delete(recordID int64) error {
	_, err := r.db.Exec(context.Background(),
		`delete from records where id = $1`, recordID)
	if err != nil {
		return errors.Wrap(err, "fail to delete record from db")
	}

	return nil
}

func (r Record) Update(recordID int64, value float32) error {
	_, err := r.db.Exec(context.Background(),
		`update records set value = $1 where id = $2`, value, recordID)
	if err != nil {
		return errors.Wrap(err, "fail to update record")
	}

	return nil
}

func (r Record) GetRecordsForDay(date time.Time) ([]entity.Record, error) {
	var records []entity.Record
	err := pgxscan.Select(context.Background(), r.db, &records,
		"select * from records where date_trunc('day', datetime) = $1", date)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get records for the day")
	}
	return records, nil
}

func (r Record) GetStatForMonth(year, month int) ([]entity.Value, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT
					date_trunc('day', datetime) as day,
    				sum(case when is_income then value else 0 end) AS incomes,
    				sum(case when not is_income then value else 0 end) AS costs
				FROM records
				WHERE date_trunc('month', datetime) = $1
				GROUP BY date_trunc('day', datetime)`, time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC))

	if err != nil {
		return nil, errors.Wrap(err, "fail to select for month")
	}
	defer rows.Close()

	var values []entity.Value
	for rows.Next() {
		var value entity.Value

		err = rows.Scan(&value.DateTime, &value.Incomes, &value.Costs)
		if err != nil {
			return nil, errors.Wrap(err, "fail to scan row in select for month")
		}
		values = append(values, value)
	}

	return values, nil
}

func (r Record) GetStatForYear(year int) ([]entity.Value, error) {
	rows, err := r.db.Query(context.Background(),
		`SELECT
					date_trunc('month', datetime) as month,
    				sum(case when is_income then value else 0 end) AS incomes,
    				sum(case when not is_income then value else 0 end) AS costs
				FROM records
				WHERE date_trunc('year', datetime) = $1
				GROUP BY date_trunc('month', datetime)
				`, time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		return nil, errors.Wrap(err, "fail to select for year")
	}
	defer rows.Close()

	var values []entity.Value
	for rows.Next() {
		var value entity.Value

		err = rows.Scan(&value.DateTime, &value.Incomes, &value.Costs)
		if err != nil {
			return nil, errors.Wrap(err, "fail to scan row in select for year")
		}

		values = append(values, value)
	}

	return values, nil
}

func (r Record) SaveBatch(records []entity.Record) error {
	batch := pgx.Batch{}
	for _, record := range records {
		batch.Queue(`INSERT INTO records(customer_id, event_id, is_income, value, datetime, description)
							VALUES ($1, $2, $3, $4, $5, $6)`,
			record.CustomerID, record.EventID, record.IsIncome,
			record.Value, record.DateTime, record.Description)
	}
	_, err := r.db.SendBatch(context.Background(), &batch).Exec()
	if err != nil {
		return errors.Wrap(err, "fail to batch save records")
	}

	return nil
}

func (r Record) GetProfit(start, end time.Time) (float32, error) {
	var result float32
	err := r.db.QueryRow(context.Background(),
		`SELECT sum(case when is_income then value else -value end)
				FROM records
				WHERE $1 >= datetime and datetime < $2`, start, end).Scan(&result)
	if err != nil {
		return 0, errors.Wrap(err, "fail to get profit")
	}

	return result, nil
}

func NewRecord(db *pgxpool.Pool) *Record {
	return &Record{db}
}

var _ repository.Record = &Record{}
