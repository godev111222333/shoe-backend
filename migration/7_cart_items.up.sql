create table cart_items
(
    id         serial primary key,
    cart_id    bigint unsigned not null,
    product_id bigint unsigned not null,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    constraint FK_cart_items_cart FOREIGN KEY (cart_id) references carts (id),
    constraint FK_cart_items_product FOREIGN KEY (product_id) references products (id)
)
