package postgresql

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
	"github.com/mark-by/little-busy-back/pkg/utils"
	"github.com/pkg/errors"
	"time"
)

type Event struct {
	db *pgxpool.Pool
}

func (e Event) DeleteAllForCustomer(customerID int64) error {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() { utils.EndTx(tx, err) }()
	_, err = tx.Exec(context.Background(), "delete from recurring_events where event_id in (select id from events where customer_id = $1)", customerID)
	if err != nil {
		return errors.Wrap(err, "fail to delete recurring events for customer")
	}

	_, err = tx.Exec(context.Background(), "delete from events where customer_id = $1", customerID)
	if err != nil {
		return errors.Wrap(err, "fail to delete events for customer")
	}

	return nil
}

func (e Event) Get(eventID int64) (*entity.Event, error) {
	event, err := e.scanVerboseEvent(e.db.QueryRow(context.Background(),
		`select e.id, customer_id, start_time, e.end_time, description, price, period, re.end_time recurring_end_time, 
       	c.name, c.special_price_per_hour
		from events e
       	left join recurring_events re on e.id = re.event_id 
		left join customers c on customer_id = c.id	
		where e.id = $1
       `, eventID))
	if err != nil {
		return nil, errors.Wrap(err, "fail to get event by id")
	}

	if event.Period != nil {
		event.IsRecurring = true
	}

	return &event, nil
}

func (e Event) GetForCustomer(customerID int64, since time.Time, days int) ([]entity.Event, error) {
	end := since.AddDate(0, 0, days)

	rows, err := e.db.Query(context.Background(),
		`select e.id as id, customer_id, e.start_time, e.end_time, description, price, 
       	c.name, c.special_price_per_hour
		from events as e
		left join recurring_events as re on e.id = re.event_id
		join customers as c on customer_id = c.id 
		where
		    re.event_id is null 
		  	and customer_id = $1
		  	and e.start_time > $2
		  	and e.start_time < $3`, customerID, since, end)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select events")
	}

	events, err := e.scanPreviewEvents(rows)

	if err != nil {
		return nil, errors.Wrap(err, "fail to select events for month")
	}

	return e.addCustomersRecurringEvents(events, since, end, customerID)
}

func (e Event) scanPreviewEvents(rows pgx.Rows) ([]entity.Event, error) {
	var events []entity.Event

	defer rows.Close()

	for rows.Next() {
		event, err := e.scanPreviewEvent(rows)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (e Event) GetForMonth(year, month int) ([]entity.Event, error) {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
	end := start.AddDate(0, 1, 0)

	rows, err := e.db.Query(context.Background(),
		`select e.id as id, customer_id, e.start_time, e.end_time, description, price, 
       	c.name, c.special_price_per_hour
		from events as e
		left join recurring_events as re on e.id = re.event_id
		left join customers as c on customer_id = c.id 
		where re.event_id is null and e.start_time > $1 and e.start_time < $2`, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select events")
	}

	events, err := e.scanPreviewEvents(rows)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select events for month")
	}

	events, err = e.addRecurringEvents(events, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "fail to add recurring events")
	}

	return events, nil
}

func (e Event) GetForDay(year, month, day int) ([]entity.Event, error) {
	start := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location())
	end := start.AddDate(0, 0, 1).Add(-1 * time.Second)

	var events entity.Events
	err := pgxscan.Select(context.Background(), e.db, &events,
		`select e.id, customer_id, e.start_time, e.end_time, description, price,
        c.id as "customer.id", c.name as "customer.name", c.special_price_per_hour "customer.special_price_per_hour", c.tel "customer.tel" 
		from events as e 
		left join recurring_events as re on e.id = re.event_id 
		join customers as c on customer_id = c.id
		where re.event_id is null and e.start_time > $1 and e.start_time < $2`, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "fail to select events")
	}

	events, err = e.addRecurringEvents(events, start, end)
	if err != nil {
		return nil, errors.Wrap(err, "fail to add recurring events")
	}

	return events, nil
}

func (e Event) Create(event *entity.Event) (*entity.Event, error) {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() { utils.EndTx(tx, err) }()
	event, err = e.create(event, tx)
	return event, err
}

func (e Event) UpdateRegular(event *entity.Event) error {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() { utils.EndTx(tx, err) }()

	_, err = tx.Exec(context.Background(),
		"update events set customer_id = $1, start_time = $2, end_time = $3, description = $4, price = $5 where id = $6",
		event.CustomerID,
		event.StartTime,
		event.EndTime,
		event.Description,
		event.Price,
		event.ID)
	if err != nil {
		return errors.Wrap(err, "fail to update")
	}

	if event.RecurringEndTime != nil && event.RecurringEndTime.IsZero() {
		event.RecurringEndTime = nil
	}

	_, err = tx.Exec(context.Background(),
		"update recurring_events set period = $1 where event_id = $2", event.Period, event.ID)
	if err != nil {
		return errors.Wrap(err, "fail to update recurring event")
	}

	if event.IsRecurring {
		err = e.makeEventRecurring(event, tx)
		if err != nil {
			return errors.Wrap(err, "fail to make event recurring")
		}
	}

	return nil
}

func (e Event) UpdateWithNextRecurring(event *entity.Event, currStartTime time.Time) error {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() { utils.EndTx(tx, err) }()

	err = e.stopRecurring(tx, event.ID, currStartTime.AddDate(0, 0, -1))
	if err != nil {
		return errors.Wrap(err, "fail to stop recurring event")
	}

	event.IsRecurring = true
	event, err = e.create(event, tx)
	if err != nil {
		return errors.Wrap(err, "fail to create new recurring")
	}

	return nil
}

func (e Event) UpdateOnlyCurrRecurring(event *entity.Event, currStartTime time.Time) error {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() { utils.EndTx(tx, err) }()

	err = e.deleteOnlyCurrRecurring(event, currStartTime, tx)
	if err != nil {
		return errors.Wrap(err, "fail to delete only curr recurring")
	}

	event.IsRecurring = false
	_, err = e.create(event, tx)
	if err != nil {
		return errors.Wrap(err, "fail to create updated event")
	}

	return nil
}

func (e Event) DeleteWithNextRecurring(event *entity.Event, currStartTime time.Time) error {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() { utils.EndTx(tx, err) }()

	if event.RecurringEndTime != nil && event.StartTime.Sub(*event.RecurringEndTime) == 0 {
		return e.deleteAllForEvent(tx, event.ID)
	}

	if event.StartTime.Sub(currStartTime) != 0 {
		err = e.stopRecurring(tx, event.ID, currStartTime.AddDate(0, 0, -1))
		return err
	}

	_, err = tx.Exec(context.Background(), "delete from recurring_events where event_id = $1", event.ID)
	if err != nil {
		return errors.Wrap(err, "fail to delete from recurring_events")
	}

	_, err = tx.Exec(context.Background(), "delete from events where id = $1", event.ID)
	if err != nil {
		return errors.Wrap(err, "fail to delete from events")
	}
	return nil
}

func (e Event) DeleteOnlyCurrRecurring(event *entity.Event, currStartTime time.Time) error {
	tx, err := e.db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer func() { utils.EndTx(tx, err) }()

	// если от повторяющегося события осталось только одно событие
	if event.RecurringEndTime != nil && event.StartTime.Sub(*event.RecurringEndTime) == 0 {
		return e.deleteAllForEvent(tx, event.ID)
	}

	err = e.deleteOnlyCurrRecurring(event, currStartTime, tx)
	return err
}

func (e Event) DeleteRegular(eventID int64) error {
	_, err := e.db.Exec(context.Background(), "delete from events where id = $1", eventID)
	if err != nil {
		return errors.Wrap(err, "fail to delete")
	}
	return nil
}

func (e Event) deleteRegular(tx pgx.Tx, eventID int64) error {
	_, err := tx.Exec(context.Background(), "delete from events where id = $1", eventID)
	return err
}

func NewEvent(db *pgxpool.Pool) *Event {
	return &Event{db}
}

var _ repository.Events = &Event{}
