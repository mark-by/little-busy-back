package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/pkg/errors"
	"time"
)

type RecordsI interface {
	Create(record *entity.Record) error
	Delete(recordID int64) error
	Get(since int64, limit int) ([]entity.Record, error)
	GetRecordsForDay(date time.Time) ([]entity.Record, error)
	GetStatForMonth(year, month int) ([]entity.Value, error)
	GetStatForYear(year int) ([]entity.Value, error)
	GetProfit(start, end time.Time) (float32, error)
	SaveFromEvents(events []entity.Event) error
	Update(recordID int64, value float32) error
}

type Records struct {
	records  repository.Record
	settings repository.Settings
}

func NewRecords(recordsRepo repository.Record, settingsRepo repository.Settings) *Records {
	return &Records{records: recordsRepo, settings: settingsRepo}
}

func (r Records) Get(since int64, limit int) ([]entity.Record, error) {
	return r.records.Select(since, limit)
}

func (r Records) SaveFromEvents(events []entity.Event) error {
	var newRecords []entity.Record

	settings, err := r.settings.Get()
	if err != nil {
		return errors.Wrap(err, "fail to save records from events")
	}

	for _, event := range events {
		if event.Customer == nil {
			continue
		}

		price := r.calculatePriceForEvent(event, settings.DefaultPricePerHour)
		if price == 0 {
			continue
		}

		eventID := event.ID
		newRecords = append(newRecords, entity.Record{
			CustomerID:  event.CustomerID,
			EventID:     &eventID,
			IsIncome:    true,
			Value:       price,
			DateTime:    event.StartTime,
			Description: &event.Customer.Name,
		})
	}

	return r.records.SaveBatch(newRecords)
}

func (r Records) calculatePriceForEvent(event entity.Event, defaultPricePerHour int) float32 {
	if event.Price != nil {
		return float32(*event.Price)
	}

	eventDurationInHours := event.EndTime.Sub(event.StartTime).Hours()

	if event.Customer.SpecialPricePerHour != nil {
		return float32(eventDurationInHours) * float32(*event.Customer.SpecialPricePerHour)
	}

	return float32(eventDurationInHours) * float32(defaultPricePerHour)
}

func (r Records) Delete(recordID int64) error {
	return r.records.Delete(recordID)
}

func (r Records) Update(recordID int64, value float32) error {
	return r.records.Update(recordID, value)
}

func (r Records) GetStatForMonth(year, month int) ([]entity.Value, error) {
	return r.records.GetStatForMonth(year, month)
}

func (r Records) GetStatForYear(year int) ([]entity.Value, error) {
	return r.records.GetStatForYear(year)
}

func (r Records) GetProfit(start, end time.Time) (float32, error) {
	return r.records.GetProfit(start, end)
}

func (r Records) Create(record *entity.Record) error {
	return r.records.Create(record)
}

func (r Records) GetRecordsForDay(date time.Time) ([]entity.Record, error) {
	return r.records.GetRecordsForDay(date)
}

var _ RecordsI = &Records{}
