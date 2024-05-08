CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(26) PRIMARY KEY,
    name VARCHAR(50),
    phone_number VARCHAR(16),
    role VARCHAR(255),
    password VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_phone_number_role UNIQUE (phone_number, role)
);

CREATE UNIQUE INDEX unique_phone_number_role_idx ON users (phone_number, role);