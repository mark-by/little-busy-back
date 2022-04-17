package postgresql

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/mark-by/little-busy-back/scheduler/internal/domain/entity"
	"github.com/pkg/errors"
	"time"
)

func (e Event) create(event *entity.Event, tx pgx.Tx) (*entity.Event, error) {
	err := tx.QueryRow(context.Background(),
		"insert into events (customer_id, start_time, end_time, description) "+
			"values ($1, $2, $3, $4) returning id",
		event.CustomerID,
		event.StartTime,
		event.EndTime,
		event.Description).Scan(&event.ID)
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
	sqlString := "select e.id, customer_id, e.start_time, " +
		"e.end_time, description, period, re.end_time as recurring_end_time " +
		"from recurring_events re join events e on re.event_id = e.id " +
		"where (e.start_time <= $1 and (re.end_time >= $2 or re.end_time is null)) "
	args = append(args, end, start)
	if forCustomer != 0 {
		sqlString += "and customer_id = $3"
		args = append(args, forCustomer)
	}

	var recurringEvents entity.Events

	err := pgxscan.Select(context.Background(), e.db, &recurringEvents, sqlString, args...)

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

func (e Event) deleteOnlyCurrRecurring(event *entity.Event, date time.Time, tx pgx.Tx) error {
	err := e.stopRecurring(tx, event.ID, date.AddDate(0, 0, -1))
	if err != nil {
		return errors.Wrap(err, "fail to stop recurring")
	}

	oldEvent, err := e.Get(event.ID)
	if err != nil {
		return errors.Wrap(err, "fail to get old event")
	}

	_, err = e.create(oldEvent.CopyWithNewDate(date).NextRecurring(), tx)
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
