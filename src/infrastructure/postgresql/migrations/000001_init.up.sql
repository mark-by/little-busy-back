CREATE TABLE users (
    id serial primary key,
    username varchar(16) not null unique,
    password varchar(32) not null,
    is_admin boolean
);

CREATE TABLE customer (
    id serial primary key,
    name varchar(32) not null unique ,
    tel varchar(10)
);

CREATE TABLE session (
    id serial primary key,
    datetime timestamp not null,
    duration smallint not null,
    is_active boolean,
    is_constant boolean,
    customer_id serial not null references customer(id) on delete cascade
);