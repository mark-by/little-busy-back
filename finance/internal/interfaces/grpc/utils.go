package grpc

import (
	"github.com/mark-by/little-busy-back/finance/internal/domain/entity"
	"github.com/mark-by/little-busy-back/finance/pkg/proto/finance"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"log"
	"time"
)

func convertRecord(record *entity.Record) *finance.Record {
	return &finance.Record{
		ClientID:    utils.DropNil(record.ClientID),
		EventID:     utils.DropNil(record.EventID),
		Type:        record.Type,
		Value:       record.Value,
		DateTime:    record.DateTime.Format(time.RFC3339),
		Description: utils.DropNil(record.Description),
	}
}

func convertRecords(records []entity.Record) *finance.Records {
	newRecords := make([]*finance.Record, 0, len(records))
	for _, record := range records {
		newRecords = append(newRecords, convertRecord(&record))
	}

	return &finance.Records{Records: newRecords}
}

func convertProtoRecord(record *finance.Record) *entity.Record {
	dateTime, err := time.Parse(time.RFC3339, record.DateTime)
	if err != nil {
		log.Printf("fail to convert proto record dateTime")
		return nil
	}
	return &entity.Record{
		ClientID:    &record.ClientID,
		EventID:     &record.EventID,
		Type:        record.Type,
		Value:       record.Value,
		DateTime:    dateTime,
		Description: &record.Description,
	}
}

func convertValues(values []entity.Value) *finance.Values {
	var newValues finance.Values
	for _, value := range values {
		newValues.Values = append(newValues.Values, &finance.Value{
			Income:   value.Incomes,
			Cost:     value.Costs,
			DateTime: value.Date.Format(time.RFC3339),
		})
	}
	return &newValues
}
