CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    transaction_id VARCHAR(100) UNIQUE NOT NULL, -- External transaction ID from payment provider
    order_id INTEGER NOT NULL REFERENCES orders(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    
    -- Payment details
    payment_method VARCHAR(50) NOT NULL CHECK (payment_method IN ('mpesa', 'card', 'bank_transfer', 'cash')),
    amount DECIMAL(10,2) NOT NULL CHECK (amount > 0),
    currency VARCHAR(3) DEFAULT 'KES',
    
    -- Transaction status
    status VARCHAR(50) DEFAULT 'pending' CHECK (status IN ('pending', 'processing', 'completed', 'failed', 'cancelled', 'refunded')),
    
    -- M-Pesa specific fields
    mpesa_receipt_number VARCHAR(100),
    mpesa_transaction_date TIMESTAMP,
    mpesa_phone_number VARCHAR(20),
    mpesa_checkout_request_id VARCHAR(100),
    
    -- General payment fields
    payment_reference VARCHAR(100), -- Can be used for any payment method
    gateway_response JSONB, -- Store full response from payment gateway
    
    -- Failure information
    failure_reason TEXT,
    failure_code VARCHAR(50),
    
    -- Refund information
    refunded_amount DECIMAL(10,2) DEFAULT 0,
    refunded_at TIMESTAMP,
    refund_reason TEXT,
    
    -- Timestamps
    initiated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create payment attempts table (for retry logic)
CREATE TABLE payment_attempts (
    id SERIAL PRIMARY KEY,
    transaction_id INTEGER NOT NULL REFERENCES transactions(id),
    attempt_number INTEGER NOT NULL,
    status VARCHAR(50) NOT NULL,
    response_data JSONB,
    error_message TEXT,
    attempted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX idx_transactions_transaction_id ON transactions(transaction_id);
CREATE INDEX idx_transactions_order_id ON transactions(order_id);
CREATE INDEX idx_transactions_user_id ON transactions(user_id);
CREATE INDEX idx_transactions_status ON transactions(status);
CREATE INDEX idx_transactions_payment_method ON transactions(payment_method);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
CREATE INDEX idx_transactions_mpesa_checkout_request_id ON transactions(mpesa_checkout_request_id);

CREATE INDEX idx_payment_attempts_transaction_id ON payment_attempts(transaction_id);

-- Create trigger to automatically update updated_at for transactions
CREATE TRIGGER update_transactions_updated_at 
    BEFORE UPDATE ON transactions 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Create function to generate transaction IDs
CREATE OR REPLACE FUNCTION generate_transaction_id() RETURNS TEXT AS $
DECLARE
    trans_id TEXT;
BEGIN
    trans_id := 'TXN' || TO_CHAR(CURRENT_TIMESTAMP, 'YYYYMMDDHH24MISS') || LPAD(FLOOR(RANDOM() * 1000)::TEXT, 3, '0');
    RETURN trans_id;
END;
$ LANGUAGE plpgsql;