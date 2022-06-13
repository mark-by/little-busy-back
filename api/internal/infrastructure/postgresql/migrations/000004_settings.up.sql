create table settings (
    start_work_hour integer not null,
    end_work_hour integer not null,
    default_price_per_hour integer not null
);

insert into settings (start_work_hour, end_work_hour, default_price_per_hour) values (8, 22, 1000);