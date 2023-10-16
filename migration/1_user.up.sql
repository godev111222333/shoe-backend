CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    phone      varchar(20) unique,
    name       varchar(255),
    birthdate  timestamp DEFAULT CURRENT_TIMESTAMP,
    avatar_url varchar(1023),
    email      varchar(255) unique,
    balance    int not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
)
