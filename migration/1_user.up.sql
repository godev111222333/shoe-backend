CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    username   varchar(255) NOT NULL,
    password   varchar(255) NOT NULL,
    name       varchar(255),
    birthdate  timestamp,
    avatar_url varchar(1023),
    phone      varchar(20),
    email      varchar(255),
    balance    int          not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
)
