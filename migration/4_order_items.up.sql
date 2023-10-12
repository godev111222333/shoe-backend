create table order_items
(
    id         serial primary key,
    order_id   bigint unsigned not null,
    product_id bigint unsigned not null,
    at_price   int not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    constraint FK_order_items_order FOREIGN KEY (order_id) references orders (id),
    constraint FK_order_items_product FOREIGN KEY (product_id) references products (id)
)
