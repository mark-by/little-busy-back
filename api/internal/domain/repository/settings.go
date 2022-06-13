package repository

import "github.com/mark-by/little-busy-back/api/internal/domain/entity"

type Settings interface {
	Get() (*entity.Settings, error)
	Update(settings *entity.Settings) error
}
