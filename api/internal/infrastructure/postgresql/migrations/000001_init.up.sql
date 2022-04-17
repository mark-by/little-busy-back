create table users (
    id serial primary key,
    username varchar(32) not null,
    password varchar(60) not null
);

create table customers (
    id serial primary key,
    name varchar(32) not null,
    tel varchar(10) unique
);