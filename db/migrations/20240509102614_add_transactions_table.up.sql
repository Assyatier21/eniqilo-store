CREATE TABLE IF NOT EXISTS transactions (
    id VARCHAR(255) PRIMARY KEY,
    customer_id VARCHAR(26),
    product_details JSONB,
    paid FLOAT,
    change FLOAT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_transaction_customer_id ON transactions (customer_id);
