CREATE TABLE Product (
    id INT UNSIGNED NOT NULL PRIMARY KEY,     -- 产品ID，非负整数
    name VARCHAR(255) NOT NULL,              -- 产品名称
    description TEXT,                        -- 产品描述
    picture VARCHAR(255),                    -- 产品图片路径或URL
    price FLOAT NOT NULL,                    -- 产品价格
    categories JSON,                          -- 产品类别，存储为JSON格式
    PRIMARY KEY (id)
);