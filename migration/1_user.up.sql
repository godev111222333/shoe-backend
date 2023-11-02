CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    phone      varchar(20) unique,
    name       varchar(255),
    avatar_url varchar(1023),
    email      varchar(255) unique,
    balance    int not null,
    status     varchar(255),
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
)
