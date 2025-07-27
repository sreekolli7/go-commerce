

## Project Overview

I built this to learn about microservices architecture, REST APIs, and database design. The project consists of 4 main services:

### Services
1. **Product Service** (Port 8080) - Handles product catalog and inventory
2. **User Service** (Port 8081) - User registration, authentication, and profiles  
3. **Checkout Service** (Port 8082) - Order processing and payment handling
4. **Scraper Service** (Port 8083) - Collects product data from external sources

## Architecture

I used Clean Architecture principles with these layers:
- **Domain** - Business entities and rules
- **Usecase** - Business logic and application services
- **Interface** - HTTP handlers and API endpoints
- **Infrastructure** - Database connections and external services

## Database Setup

You need PostgreSQL running with a database called `gocommerce`. The connection string is:
```
postgres://postgres:postgres123@localhost:5432/gocommerce?sslmode=disable
```

## How to Run

### 1. Start Product Service
```bash
cd product-service/cmd
go run main.go
```

### 2. Start User Service  
```bash
cd user-service/cmd
go run main.go
```

### 3. Start Checkout Service
```bash
cd checkout-service/cmd
go run main.go
```

### 4. Start Scraper Service
```bash
cd checkout-service/scraper-service/cmd
go run main.go
```

## API Endpoints

### Product Service (Port 8080)
- `GET /products` - List all products
- `POST /products` - Create new product
- `GET /products/:id` - Get product by ID

### User Service (Port 8081)
- `POST /users/register` - Register new user
- `POST /users/login` - User login
- `GET /users/:id` - Get user profile

### Checkout Service (Port 8082)
- `POST /checkout/` - Create new order
- `GET /checkout/orders/:user_id` - Get orders for user

### Scraper Service (Port 8083)
- `POST /scrape` - Scrape product from URL
- `GET /products` - List scraped products
- `GET /products/:external_id` - Get scraped product

## Testing

You can test the APIs using curl:

```bash
# Create an order
curl -X POST http://localhost:8082/checkout/ \
  -H "Content-Type: application/json" \
  -d '{"user_id":1,"items":[{"product_id":101,"quantity":2,"price":19.99}],"payment_method":"credit_card"}'

# Get orders for user
curl http://localhost:8082/checkout/orders/1
```

## Technologies Used

- **Go** - Programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **sqlx** - Database driver
- **JWT** - Authentication tokens

## What I Learned

- Microservices architecture and communication
- RESTful API design
- Database design and SQL
- Clean Architecture principles
- Go programming and concurrency
- Docker containerization (basic)

## Future Improvements

- Add authentication middleware
- Implement real payment processing
- Add more comprehensive error handling
- Add unit tests
- Deploy to cloud platform
