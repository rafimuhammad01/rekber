CREATE TABLE IF NOT EXISTS users(
   id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
   name VARCHAR(50) NOT NULL,
   phone_number VARCHAR(50) UNIQUE NOT NULL,
   phone_number_verified_at TIMESTAMP DEFAULT NULL,
   created_at TIMESTAMP DEFAULT NOW()
);