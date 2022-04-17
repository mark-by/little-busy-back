package clickhouse

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/mark-by/little-busy-back/finance/internal/domain/entity"
	"github.com/mark-by/little-busy-back/finance/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type Value struct {
	db clickhouse.Conn
}

func (v Value) GetForMonth(year, month int) ([]entity.Value, error) {
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

func (v Value) GetForYear(year int) ([]entity.Value, error) {
	rows, err := v.db.Query(context.Background(),
		`SELECT
					toMonth(DateTime) as month,
    				sum(if(Type = 'income', Value, 0)) AS incomes,
    				sum(if(Type = 'cost', Value, 0)) AS costs
				FROM Records
				GROUP BY toMonth(DateTime)
				HAVING toYear(DateTime) = $1`, year)
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

		value.Date = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		values = append(values, value)
	}

	return values, nil
}

func NewValue(db clickhouse.Conn) *Value {
	return &Value{db}
}

var _ repository.Value = &Value{}
