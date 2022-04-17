package microservices

import (
	"context"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/mark-by/little-busy-back/finance/pkg/proto/finance"
	"github.com/mark-by/little-busy-back/pkg/converter"
)

type Value struct {
	financeClient finance.FinanceServiceClient
}

func (v Value) ProfitForMonth(year, month int) (float32, error) {
	profit, err := v.financeClient.ProfitForMonth(context.Background(), &finance.DateRequest{
		Year:  int32(year),
		Month: int32(month),
	})

	if err != nil {
		return 0, err
	}

	return profit.Value, nil
}

func (v Value) GetForMonth(year, month int) ([]entity.Value, error) {
	values, err := v.financeClient.GetValuesForMonth(context.Background(), &finance.DateRequest{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		return nil, err
	}

	return converter.ConvertPointerSlice[finance.Value, entity.Value](values.Values)
}

func (v Value) GetForYear(year int) ([]entity.Value, error) {
	values, err := v.financeClient.GetValuesForYear(context.Background(), &finance.DateRequest{
		Year: int32(year),
	})
	if err != nil {
		return nil, err
	}

	return converter.ConvertPointerSlice[finance.Value, entity.Value](values.Values)
}

func NewValue(client finance.FinanceServiceClient) *Value {
	return &Value{financeClient: client}
}

var _ repository.Value = &Value{}
