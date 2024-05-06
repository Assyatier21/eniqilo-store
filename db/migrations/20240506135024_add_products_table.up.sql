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
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);