CREATE TABLE customer (
    username text not null,
    chat_id serial not null unique,
    tel varchar(10) not null,
    notification_is_enabled bool not null default false
);