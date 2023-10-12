create table employees
(
    id         serial primary key,
    username   varchar(255) not null,
    password   varchar(255) not null,
    name       varchar(255) not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
)
