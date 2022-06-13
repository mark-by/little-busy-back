create table records (
    id serial primary key,
    customer_id serial references customers(id) on delete cascade,
    event_id serial references events(id) on delete cascade,
    is_income boolean,
    value double precision not null,
    datetime timestamp with time zone not null,
    description text
);
