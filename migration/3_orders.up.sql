create table orders
(
    id             serial primary key,
    user_id        bigint unsigned not null,
    payment_status int,
    created_at     timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at     timestamp DEFAULT CURRENT_TIMESTAMP,
    constraint FK_user FOREIGN KEY (user_id) references users (id)
)
