create table carts
(
    id             serial primary key,
    user_id        bigint unsigned not null,
    created_at     timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at     timestamp DEFAULT CURRENT_TIMESTAMP,
    constraint FK_carts_user FOREIGN KEY (user_id) references users (id)
)
