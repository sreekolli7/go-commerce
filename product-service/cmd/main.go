package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sreekolli7/go-commerce/product-service/internal/infrastructure"
	_interface "github.com/sreekolli7/go-commerce/product-service/internal/interface"
	"github.com/sreekolli7/go-commerce/product-service/internal/usecase"
)

func main() {
	// Get the database URL from environment variable or use a default
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres123@localhost:5432/gocommerce?sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up repository, usecase, and handler
	repo := infrastructure.NewPostgresProductRepo(db)
	uc := &usecase.ProductUsecase{Repo: repo}
	handler := &_interface.ProductHandler{Usecase: uc}

	// Set up Gin router and register routes
	r := gin.Default()
	handler.RegisterRoutes(r)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
