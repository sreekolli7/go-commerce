# Project Structure

This document shows how I organized my e-commerce microservices project.

```
gocommerce/
├── README.md                    # Project documentation
├── run_services.sh             # Script to start all services
├── setup_database.sql          # Database setup script
├── PROJECT_STRUCTURE.md        # This file
│
├── product-service/            # Product catalog service
│   ├── cmd/
│   │   └── main.go           # Entry point
│   ├── internal/
│   │   ├── domain/           # Business entities
│   │   ├── usecase/          # Business logic
│   │   ├── interface/        # HTTP handlers
│   │   └── infrastructure/   # Database layer
│   └── go.mod
│
├── user-service/              # User management service
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── interface/
│   │   └── infrastructure/
│   └── go.mod
│
├── checkout-service/          # Order processing service
│   ├── cmd/
│   │   └── main.go
│   ├── internal/
│   │   ├── domain/
│   │   ├── usecase/
│   │   ├── interface/
│   │   └── infrastructure/
│   ├── scraper-service/      # Product scraping service
│   │   ├── cmd/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── domain/
│   │   │   ├── usecase/
│   │   │   ├── interface/
│   │   │   └── infrastructure/
│   │   └── go.mod
│   └── go.mod
│
└── .git/                     # Git repository
```

## Architecture Explanation

I used **Clean Architecture** principles for each service:

1. **Domain Layer** - Contains business entities and rules
2. **Usecase Layer** - Contains business logic and application services  
3. **Interface Layer** - Contains HTTP handlers and API endpoints
4. **Infrastructure Layer** - Contains database connections and external services


## Why This Structure?

I chose this structure because:
- It separates concerns clearly
- Each service is independent
- Easy to add new features
- Follows Go best practices
- Makes testing easier

## Database Design

I used PostgreSQL with these main tables:
- `users` - User accounts and authentication
- `products` - Product catalog
- `orders` - Order information
- `scraped_products` - Products from external sources

## Communication Between Services

Currently, services communicate through:
- REST APIs
- Shared database (for simplicity)
- JSON over HTTP
