package postgresql

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/pkg/errors"
	"time"
)

func (e Event) create(event *entity.Event, tx pgx.Tx) (*entity.Event, error) {
	err := tx.QueryRow(context.Background(),
		"insert into events (customer_id, start_time, end_time, description, price) "+
			"values ($1, $2, $3, $4, $5) returning id",
		event.CustomerID,
		event.StartTime,
		event.EndTime,
		event.Description,
		event.Price).Scan(&event.ID)
	if err != nil {
		return nil, errors.Wrap(err, "fail to insert")
	}

	if event.IsRecurring {
		err = e.makeEventRecurring(event, tx)
		if err != nil {
			return nil, errors.Wrap(err, "fail to create recurring event")
		}
	}

	return event, nil
}

func (e Event) makeEventRecurring(event *entity.Event, tx pgx.Tx) error {
	if event.RecurringEndTime != nil && event.RecurringEndTime.IsZero() {
		event.RecurringEndTime = nil
	}
	_, err := tx.Exec(context.Background(),
		"insert into recurring_events (event_id, week_day, day, period, end_time) "+
			"values ($1, $2, $3, $4, $5)",
		event.ID,
		int(event.StartTime.Weekday()),
		event.StartTime.Day(),
		event.Period,
		event.RecurringEndTime)
	return err
}

func (e Event) getRecurringEvents(start, end time.Time, forCustomer int64) (entity.Events, error) {
	var args []interface{}
	sqlString := `select e.id, customer_id, e.start_time, 
       	e.end_time, price, description, period, re.end_time as recurring_end_time, c.name
		from recurring_events re
		join events e on re.event_id = e.id 
		left join customers c on customer_id = c.id 
		where (e.start_time <= $1 and (re.end_time >= $2 or re.end_time is null)) `
	args = append(args, end, start)
	if forCustomer != 0 {
		sqlString += "and customer_id = $3"
		args = append(args, forCustomer)
	}

	var recurringEvents entity.Events

	rows, err := e.db.Query(context.Background(), sqlString, args...)
	if err != nil {
		return nil, errors.Wrap(err, "fail to query recurring events:")
	}
	defer rows.Close()

	for rows.Next() {
		var newEvent entity.Event

		customerID := sql.NullInt64{}
		customerName := sql.NullString{}
		price := sql.NullFloat64{}
		description := sql.NullString{}
		period := sql.NullString{}
		recurringEndTime := sql.NullTime{}

		err = rows.Scan(&newEvent.ID, &customerID, &newEvent.StartTime, &newEvent.EndTime, &price, &description, &period,
			&recurringEndTime, &customerName)
		if err != nil {
			return nil, errors.Wrap(err, "fail to scan recurring event:")
		}

		if customerID.Valid {
			newEvent.CustomerID = &customerID.Int64
			newEvent.Customer = &entity.Customer{
				ID:   customerID.Int64,
				Name: customerName.String,
			}
		}
		if price.Valid {
			newEvent.Price = &price.Float64
		}
		if description.Valid {
			newEvent.Description = &description.String
		}
		if period.Valid {
			newEvent.Period = &period.String
		}
		if recurringEndTime.Valid {
			newEvent.RecurringEndTime = &recurringEndTime.Time
		}

		recurringEvents = append(recurringEvents, newEvent)
	}

	if err != nil {
		return nil, errors.Wrap(err, "fail to select recurring events")
	}

	return recurringEvents, nil
}

func (e Event) addRecurringEvents(events entity.Events, start, end time.Time) (entity.Events, error) {
	return e.addCustomersRecurringEvents(events, start, end, 0)
}

func (e Event) addCustomersRecurringEvents(events entity.Events, start, end time.Time, customerID int64) (entity.Events, error) {
	recurringEvents, err := e.getRecurringEvents(start, end, customerID)
	if err != nil {
		return nil, err
	}

	return newRecurringEventsAdder(recurringEvents).AddTo(events, start, end), nil
}

func (e Event) deleteRecurringEvent(tx pgx.Tx, eventID int64) error {
	_, err := tx.Exec(context.Background(), "delete from recurring_events where event_id = $1", eventID)
	return err
}

func (e Event) deleteAllForEvent(tx pgx.Tx, eventID int64) error {
	err := e.deleteRecurringEvent(tx, eventID)
	if err != nil {
		return errors.Wrap(err, "fail to delete regular event")
	}
	err = e.deleteRegular(tx, eventID)
	if err != nil {
		return errors.Wrap(err, "fail to recurring event")
	}
	return nil
}

// deleteOnlyCurrRecurring предназначена для удаления конкретного повторяющегося события
// функция делит одно повторяющееся событие на два. До текущего и после. Делает пометку для события до, что повторение
// события останавливается днем date - 1. И создает копию события после
func (e Event) deleteOnlyCurrRecurring(event *entity.Event, date time.Time, tx pgx.Tx) error {
	oldEvent, err := e.Get(event.ID)
	if err != nil {
		return errors.Wrap(err, "fail to get old event")
	}

	err = e.stopRecurring(tx, event.ID, date.AddDate(0, 0, -1))
	if err != nil {
		return errors.Wrap(err, "fail to stop recurring")
	}

	nextEvent := oldEvent.CopyWithNewDate(date).NextRecurring()
	if nextEvent == nil {
		return nil
	}

	_, err = e.create(nextEvent, tx)
	if err != nil {
		return errors.Wrap(err, "fail to create next recurring")
	}
	return nil
}

func (e Event) stopRecurring(tx pgx.Tx, eventID int64, endTime time.Time) error {
	_, err := tx.Exec(context.Background(),
		"update recurring_events set end_time = $1 where event_id = $2",
		endTime, eventID)
	return err
}

type canScan interface {
	Scan(dest ...interface{}) error
}

func (e Event) scanPreviewEvent(row canScan) (entity.Event, error) {
	newEvent := entity.Event{}
	customer := entity.Customer{}
	customerID := sql.NullInt64{}
	description := sql.NullString{}
	price := sql.NullFloat64{}

	customerName := sql.NullString{}
	customerPrice := sql.NullInt32{}

	err := row.Scan(&newEvent.ID, &customerID, &newEvent.StartTime, &newEvent.EndTime, &description, &price,
		&customerName, &customerPrice)

	if err != nil {
		return newEvent, errors.Wrap(err, "fail to scan for select events")
	}

	if customerID.Valid {
		newEvent.CustomerID = &customerID.Int64
		customer.Name = customerName.String
		if customerPrice.Valid {
			tmp := int(customerPrice.Int32)
			customer.SpecialPricePerHour = &tmp
		}
		newEvent.Customer = &customer
	}

	return newEvent, nil
}

func (e Event) scanVerboseEvent(row canScan) (entity.Event, error) {
	newEvent := entity.Event{}
	customer := entity.Customer{}
	customerID := sql.NullInt64{}
	description := sql.NullString{}
	price := sql.NullFloat64{}
	period := sql.NullString{}
	recurringEndTime := sql.NullTime{}

	customerName := sql.NullString{}
	customerPrice := sql.NullInt32{}

	err := row.Scan(&newEvent.ID, &customerID, &newEvent.StartTime, &newEvent.EndTime, &description, &price,
		&period, &recurringEndTime, &customerName, &customerPrice)

	if err != nil {
		return newEvent, errors.Wrap(err, "fail to scan for select events")
	}

	if price.Valid {
		newEvent.Price = &price.Float64
	}

	if customerID.Valid {
		newEvent.CustomerID = &customerID.Int64
		customer.ID = customerID.Int64
		customer.Name = customerName.String
		if customerPrice.Valid {
			tmp := int(customerPrice.Int32)
			customer.SpecialPricePerHour = &tmp
		}
		newEvent.Customer = &customer
	}

	if period.Valid {
		newEvent.Period = &period.String
	}

	if recurringEndTime.Valid {
		newEvent.RecurringEndTime = &recurringEndTime.Time
	}

	return newEvent, nil
}
