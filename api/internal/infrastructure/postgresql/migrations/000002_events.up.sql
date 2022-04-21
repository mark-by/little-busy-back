create table events (
    id serial primary key,
    customer_id serial references customers(id) on delete cascade,
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    price double precision not null default 0,
    description text
);

create type period as enum ('daily', 'weekly', 'monthly');

create table recurring_events (
    event_id serial primary key references events(id) on delete cascade,
    end_time timestamp with time zone,
    week_day smallint,
    day smallint,
    period period not null default 'weekly'
);
