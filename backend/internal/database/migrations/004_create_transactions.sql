-- 004_create_transactions.sql
CREATE TYPE transaction_status AS ENUM ('pending', 'completed', 'failed', 'cancelled');
CREATE TYPE payment_method AS ENUM ('mpesa', 'card', 'cash');

CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    transaction_id VARCHAR(100) UNIQUE NOT NULL,
    mpesa_receipt_number VARCHAR(50),
    payment_method payment_method NOT NULL,
    amount DECIMAL(10,2) NOT NULL CHECK (amount >= 0),
    phone_number VARCHAR(20),
    status transaction_status DEFAULT 'pending',
    callback_data JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX idx_transactions_order ON transactions(order_id);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_method ON transactions(payment_method);
CREATE INDEX idx_transactions_receipt ON transactions(mpesa_receipt_number);

