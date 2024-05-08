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

-- Index for ULID
CREATE INDEX idx_products_id ON products (id);

-- Index for filtering by name
CREATE INDEX idx_products_name ON products (name);

-- Index for filtering by category
CREATE INDEX idx_products_category ON products (category);

-- Index for filtering by SKU
CREATE INDEX idx_products_sku ON products (sku);


-- Index for filtering by availability
CREATE INDEX idx_products_is_available ON products (is_available);


