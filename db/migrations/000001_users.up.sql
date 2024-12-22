CREATE TABLE users (
    id SERIAL PRIMARY KEY,           -- Auto-incrementing unique identifier
    username VARCHAR(50) NOT NULL,  -- Username, up to 50 characters, required
    password VARCHAR(255) NOT NULL, -- Password, stored as a hash, up to 255 characters, required
    primary_device VARCHAR(255) NOT NULL, -- Primary device ID, required
    sex CHAR(1) CHECK (sex IN ('M', 'F', 'O')), -- Sex: M for Male, F for Female, O for Other
    age INT CHECK (age >= 0),        -- Age, must be a non-negative integer
    is_deleted BOOLEAN DEFAULT FALSE -- Indicates if the user is soft-deleted
);
