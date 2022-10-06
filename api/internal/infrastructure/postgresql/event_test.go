package postgresql

import (
	"fmt"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"testing"
	"time"
)

func TestRecurringEventsAdder(t *testing.T) {
	var initialEvents []entity.Event
	period := "daily"

	end := time.Date(2022, 10, 10, 11, 0, 0, 0, time.Local)
	result := newRecurringEventsAdder([]entity.Event{
		{
			ID:               153,
			StartTime:        time.Date(2022, 10, 6, 12, 0, 0, 0, time.Local),
			EndTime:          time.Date(2022, 10, 6, 14, 0, 0, 0, time.Local),
			IsRecurring:      true,
			Period:           &period,
			RecurringEndTime: &end,
		},
	}).AddTo(initialEvents, time.Date(2021, 10, 10, 11, 0, 0, 0, time.Local),
		time.Date(2022, 10, 30, 0, 0, 0, 0, time.Local))

	for _, item := range result {
		fmt.Printf("%+v\n", item)
	}
}
