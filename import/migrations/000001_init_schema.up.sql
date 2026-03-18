create table product
(
    id          uuid                     default uuidv7() not null constraint product_pk primary key,
    name        varchar(255)                              not null,
    description text,
    sale_price  numeric(10, 2)                            not null,
    price       numeric(10, 2)                            not null,
    created_at  timestamp with time zone default now()    not null,
    updated_at  timestamp with time zone
);