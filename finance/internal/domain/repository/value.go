package repository

import "github.com/mark-by/little-busy-back/finance/internal/domain/entity"

type Value interface {
	GetForMonth(year, month int) ([]entity.Value, error)
	GetForYear(year int) ([]entity.Value, error)
}
