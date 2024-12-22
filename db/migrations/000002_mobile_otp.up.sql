CREATE TABLE mobile_otp (
    id SERIAL PRIMARY KEY,            -- Auto-incrementing unique identifier
    user_id INT NOT NULL,             -- Foreign key to reference users table
    otp VARCHAR(10) NOT NULL,         -- OTP code (e.g., 6-digit numeric or alphanumeric)
    valid_till TIMESTAMP NOT NULL,    -- Expiration time of the OTP
    created_at TIMESTAMP DEFAULT NOW() -- Time when the OTP was created
);
