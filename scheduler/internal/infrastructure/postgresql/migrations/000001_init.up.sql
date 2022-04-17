CREATE EXTENSION IF NOT EXISTS timescaledb;

create table events (
    id serial,
    customer_id serial,
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    price double precision not null default 0,
    description text
);

-- select create_hypertable('events', 'start_time');

create type period as enum ('daily', 'weekly', 'monthly');

create table recurring_events (
    event_id serial primary key,
    end_time timestamp with time zone,
    week_day smallint,
    day smallint,
    period period not null default 'weekly'
);

select exists(select from events);