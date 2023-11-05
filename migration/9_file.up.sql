create table files
(
    uuid       varchar(256) not null,
    extension  varchar(256) not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP
)
