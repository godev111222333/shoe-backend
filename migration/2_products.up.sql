create table products
(
    id          serial primary key,
    name        varchar(255)  not null,
    description varchar(1023) not null,
    price       int           not null,
    image_url   varchar(1023) not null,
    created_at  timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at  timestamp DEFAULT CURRENT_TIMESTAMP
)
