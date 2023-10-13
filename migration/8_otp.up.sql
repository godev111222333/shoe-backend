create table otp
(
    id         serial primary key,
    phone      varchar(20)  not null,
    code       varchar(10)  not null,
    status     varchar(255) not null,
    expires_at timestamp    not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    constraint FK_otp_phone FOREIGN KEY (phone) references users (phone)
)
