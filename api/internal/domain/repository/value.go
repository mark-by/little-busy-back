package repository

import "github.com/mark-by/little-busy-back/api/internal/domain/entity"

type Value interface {
	GetForMonth(year, month int) ([]entity.Value, error)
	GetForYear(year int) ([]entity.Value, error)
	ProfitForMonth(year, month int) (float32, error)
}
