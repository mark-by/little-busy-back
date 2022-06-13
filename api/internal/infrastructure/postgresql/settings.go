package postgresql

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/pkg/errors"
)

type Settings struct {
	db *pgxpool.Pool
}

func (s Settings) Get() (*entity.Settings, error) {
	var settings entity.Settings
	err := pgxscan.Get(context.Background(), s.db, &settings,
		`SELECT * FROM settings LIMIT 1`)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select settings")
	}

	return &settings, nil
}

func (s Settings) Update(settings *entity.Settings) error {
	_, err := s.db.Exec(context.Background(),
		`UPDATE settings set start_work_hour = $1, end_work_hour = $2, default_price_per_hour = $3`,
		settings.StartWorkHour,
		settings.EndWorkHour,
		settings.DefaultPricePerHour)
	if err != nil {
		return errors.Wrap(err, "fail to update settings")
	}

	return nil
}

func NewSettings(db *pgxpool.Pool) *Settings {
	return &Settings{db}
}

var _ repository.Settings = &Settings{}
