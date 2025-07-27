#!/bin/bash

# GoCommerce - Service Test Script
# This script tests if all our microservices are working correctly

echo "🧪 Testing GoCommerce Microservices..."
echo "====================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to test a service
test_service() {
    local service_name=$1
    local url=$2
    local method=$3
    local data=$4
    
    echo -n "Testing $service_name... "
    
    if [ "$method" = "GET" ]; then
        response=$(curl -s -w "%{http_code}" "$url")
    else
        response=$(curl -s -w "%{http_code}" -X "$method" -H "Content-Type: application/json" -d "$data" "$url")
    fi
    
    http_code="${response: -3}"
    body="${response%???}"
    
    if [[ "$http_code" == 2* ]]; then
        echo -e "${GREEN}✅ OK${NC}"
    else
        echo -e "${RED}❌ Failed (HTTP $http_code)${NC}"
        echo "Response: $body"
    fi
}

# Wait for services to start
echo "Waiting for services to start..."
sleep 5

# Test Checkout Service
echo ""
echo "Testing Checkout Service:"
test_service "Create Order" "http://localhost:8082/checkout/" "POST" '{"user_id":1,"items":[{"product_id":101,"quantity":2,"price":19.99}],"payment_method":"credit_card"}'
test_service "Get Orders" "http://localhost:8082/checkout/orders/1" "GET"

# Test Scraper Service
echo ""
echo "Testing Scraper Service:"
test_service "List Products" "http://localhost:8083/products" "GET"

# Test Product Service (if running)
echo ""
echo "Testing Product Service:"
test_service "Get Products" "http://localhost:8080/products" "GET"

# Test User Service (if running)
echo ""
echo "Testing User Service:"
test_service "Get Users" "http://localhost:8081/users" "GET"

echo ""
echo " Testing complete!"
echo ""
echo "If you see ✅ OK, your services are working correctly!"
echo "If you see ❌ Failed, check that the service is running and the database is set up." 