# Setup Instructions for GoCommerce Project

This document explains how to set up and run my e-commerce microservices project.

## Prerequisites

Before you start, make sure you have:
- ✅ Go 1.24+ installed
- ✅ PostgreSQL installed and running
- ✅ Git installed

## Step 1: Database Setup

1. **Start PostgreSQL** (if not already running)
2. **Create the database:**
   ```sql
   CREATE DATABASE gocommerce;
   ```
3. **Run the setup script:**
   ```bash
   psql -d gocommerce -f setup_database.sql
   ```

## Step 2: Run the Services

### Option A: Run All Services at Once
```bash
./run_services.sh
```

### Option B: Run Services Individually

1. **Product Service (Port 8080):**
   ```bash
   cd product-service/cmd
   go run main.go
   ```

2. **User Service (Port 8081):**
   ```bash
   cd user-service/cmd
   go run main.go
   ```

3. **Checkout Service (Port 8082):**
   ```bash
   cd checkout-service/cmd
   go run main.go
   ```

4. **Scraper Service (Port 8083):**
   ```bash
   cd checkout-service/scraper-service/cmd
   go run main.go
   ```

## Step 3: Test the Services

Run the test script to verify everything is working:
```bash
./test_services.sh
```

## Step 4: Manual Testing

You can also test manually using curl:

### Create an Order
```bash
curl -X POST http://localhost:8082/checkout/ \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"items":[{"product_id":101,"quantity":2,"price":19.99}],"payment_method":"credit_card"}'
```

### Get Orders for User
```bash
curl http://localhost:8082/checkout/orders/1
```

### List Products
```bash
curl http://localhost:8080/products
```

## Troubleshooting

### Common Issues:

1. **"Connection refused" error:**
   - Make sure PostgreSQL is running
   - Check the database connection string in main.go files

2. **"Port already in use" error:**
   - Stop any existing services
   - Use different ports if needed

3. **"Database doesn't exist" error:**
   - Run the database setup script
   - Check PostgreSQL is running

4. **"Module not found" error:**
   - Run `go mod tidy` in each service directory
   - Make sure all dependencies are installed

## Project Structure

- `README.md` - Main project documentation
- `run_services.sh` - Script to start all services
- `test_services.sh` - Script to test all services
- `setup_database.sql` - Database setup script
- `PROJECT_STRUCTURE.md` - Detailed project structure

## What Each Service Does

1. **Product Service** - Manages product catalog and inventory
2. **User Service** - Handles user registration and authentication
3. **Checkout Service** - Processes orders and payments
4. **Scraper Service** - Collects product data from external sources

## Next Steps

After getting the basic services running, you can:
- Add more API endpoints
- Implement authentication middleware
- Add unit tests
- Deploy to a cloud platform
- Add a frontend application
