create table records (
    customer_id serial references customers(id) on delete cascade,
    event_id serial references events(id) on delete cascade,
    is_income boolean,
    value double precision not null,
    datetime timestamp with time zone not null,
    description text
);

create table settings (
    price_per_hour double precision not null,
    start_work_hour integer,
    end_work_hour integer
)