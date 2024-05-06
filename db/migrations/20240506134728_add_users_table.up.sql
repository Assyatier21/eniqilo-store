CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(50),
    phone_number VARCHAR(16) UNIQUE,
    role VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

