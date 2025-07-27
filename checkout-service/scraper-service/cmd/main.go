package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_interface "github.com/sreekolli7/scraper-service/internal/interface"
	"github.com/sreekolli7/scraper-service/internal/usecase"
)

func main() {
	// I'm getting the database connection string from environment or using a default
	// This is how I handle configuration in my scraper service
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres123@localhost:5432/gocommerce?sslmode=disable"
	}

	// I'm connecting to the PostgreSQL database
	// This is where I'll store the scraped product data
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// I'm creating all the pieces I need for my scraper service
	// This follows the same Clean Architecture pattern I used in checkout service
	usecase := &usecase.ScraperUsecase{}             // This has my scraping business logic
	handler := _interface.NewScraperHandler(usecase) // This handles HTTP requests

	// I'm setting up my web server using Gin framework
	// This is the same framework I used for the checkout service
	r := gin.Default()

	// I'm registering my API routes
	// These are the endpoints that other services can call
	r.POST("/scrape", handler.Scrape)
	r.GET("/products/:external_id", handler.GetProduct)
	r.GET("/products", handler.ListProducts)

	// I'm starting my server on port 8083
	// I chose 8083 so it doesn't conflict with my other services
	log.Println("Scraper Service starting on port 8083...")
	if err := r.Run(":8083"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
