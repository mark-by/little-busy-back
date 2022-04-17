package application

import (
	"github.com/mark-by/little-busy-back/finance/internal/domain/entity"
	"github.com/mark-by/little-busy-back/finance/internal/domain/repository"
)

type ValueI interface {
	GetProfitForMonth(year, month int) (float32, error)
	GetForMonth(year, month int) ([]entity.Value, error)
	GetForYear(year int) ([]entity.Value, error)
}

type Value struct {
	values repository.Value
}

func (v Value) GetProfitForMonth(year, month int) (float32, error) {
	values, err := v.values.GetForMonth(year, month)
	if err != nil {
		return 0, err
	}

	var (
		incomes float32
		costs   float32
	)

	for _, value := range values {
		incomes += value.Incomes
		costs += value.Costs
	}

	return incomes - costs, nil
}

func (v Value) GetForMonth(year, month int) ([]entity.Value, error) {
	return v.values.GetForMonth(year, month)
}

func (v Value) GetForYear(year int) ([]entity.Value, error) {
	return v.values.GetForYear(year)
}

func NewValue(values repository.Value) *Value {
	return &Value{values}
}

var _ ValueI = &Value{}
