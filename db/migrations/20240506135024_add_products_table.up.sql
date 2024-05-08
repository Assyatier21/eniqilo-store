CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(30),
    sku VARCHAR(30),
    category VARCHAR(50),
    image_url VARCHAR(255),
    price FLOAT,
    stock INT,
    location VARCHAR(200),
    is_available BOOLEAN,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL
);

CREATE INDEX idx_products_id ON products (id);
CREATE INDEX idx_products_name ON products (name);
CREATE INDEX idx_products_category ON products (category);
CREATE INDEX idx_products_sku ON products (sku);


