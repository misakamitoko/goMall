CREATE TABLE shopping_carts (
    id BIGINT  AUTO_INCREMENT, -- 购物车ID
    user_id int UNSIGNED NOT NULL,              -- 用户ID（外键）
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
    PRIMARY KEY (id)
);


CREATE TABLE cart_items (
    id BIGINT AUTO_INCREMENT, -- 项目ID
    cart_id BIGINT NOT NULL,              -- 购物车ID（外键）
    product_id int UNSIGNED NOT NULL,           -- 商品ID
    quantity INT NOT NULL DEFAULT 1,      -- 商品数量
    price FLOAT NOT NULL,        -- 单价（当时价格）
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 添加时间
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
    PRIMARY KEY (id)
);


