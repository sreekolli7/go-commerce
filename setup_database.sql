-- GoCommerce Database Setup Script
-- This creates the tables needed for my e-commerce microservices project

-- Create database (run this first if database doesn't exist)
-- CREATE DATABASE gocommerce;

-- Connect to the database
-- \c gocommerce;

-- Users table (for User Service)
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Products table (for Product Service)
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock_quantity INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Orders table (for Checkout Service)
CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING',
    items JSONB, -- Store order items as JSON
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Scraped products table (for Scraper Service)
CREATE TABLE IF NOT EXISTS scraped_products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price DECIMAL(10,2),
    source VARCHAR(255),
    external_id VARCHAR(255) UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert some sample data
INSERT INTO products (name, description, price, stock_quantity) VALUES
('Laptop', 'High-performance laptop for gaming and work', 999.99, 10),
('Mouse', 'Wireless gaming mouse', 49.99, 50),
('Keyboard', 'Mechanical keyboard with RGB', 129.99, 25)
ON CONFLICT DO NOTHING;

-- Insert a sample user (password: password123)
INSERT INTO users (email, password_hash, name) VALUES
('test@example.com', '$2a$10$example_hash_here', 'Test User')
ON CONFLICT DO NOTHING;

-- Show what we created
SELECT 'Database setup complete!' as status; 