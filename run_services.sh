#!/bin/bash

# GoCommerce - Service Runner Script
# This script starts all the microservices for my e-commerce project

echo "🚀 Starting GoCommerce Microservices..."
echo "======================================"

# Function to start a service
start_service() {
    local service_name=$1
    local service_path=$2
    local port=$3
    
    echo "Starting $service_name on port $port..."
    cd "$service_path" && go run main.go &
    sleep 2
}

# Start all services
echo "1. Starting Product Service..."
start_service "Product Service" "product-service/cmd" "8080"

echo "2. Starting User Service..."
start_service "User Service" "user-service/cmd" "8081"

echo "3. Starting Checkout Service..."
start_service "Checkout Service" "checkout-service/cmd" "8082"

echo "4. Starting Scraper Service..."
start_service "Scraper Service" "checkout-service/scraper-service/cmd" "8083"

echo ""
echo "✅ All services started!"
echo ""
echo "Services running on:"
echo "- Product Service: http://localhost:8080"
echo "- User Service: http://localhost:8081" 
echo "- Checkout Service: http://localhost:8082"
echo "- Scraper Service: http://localhost:8083"
echo ""
echo "Press Ctrl+C to stop all services"

# Wait for user to stop
wait 