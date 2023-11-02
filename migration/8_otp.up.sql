create table otps
(
    id         serial primary key,
    type       varchar(20)  not null,
    email      varchar(255)  not null,
    code       varchar(10)  not null,
    status     varchar(255) not null,
    expires_at timestamp    not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
)
