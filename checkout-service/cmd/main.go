package main

import (
	"log"
	"os"

	"github.com/sreekolli7/checkout-service/internal/infrastructure"
	_interface "github.com/sreekolli7/checkout-service/internal/interface"
	"github.com/sreekolli7/checkout-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// I'm getting the database URL from environment or using a default
	// This is how I learned to handle configuration in Go
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres123@localhost:5432/gocommerce?sslmode=disable"
	}

	// I'm connecting to the database here
	// Learned this from my database class
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// I'm creating all the pieces I need for my checkout service
	// This follows the Clean Architecture pattern I learned in class
	repo := infrastructure.NewPostgresOrderRepo(db)     // This handles database stuff
	uc := &usecase.CheckoutUsecase{Repo: repo}          // This has my business logic
	handler := &_interface.CheckoutHandler{Usecase: uc} // This handles HTTP requests

	// I'm setting up my web server using Gin framework
	// Gin is really popular for Go web services
	r := gin.Default()
	handler.RegisterRoutes(r)

	// I'm starting my server on port 8082
	// Each service needs its own port so they don't conflict
	log.Println("Checkout Service running on port 8082...")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
