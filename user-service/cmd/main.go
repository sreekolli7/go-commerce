package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/sreekolli7/go-commerce/user-service/internal/infrastructure"
	_interface "github.com/sreekolli7/go-commerce/user-service/internal/interface"
	"github.com/sreekolli7/go-commerce/user-service/internal/usecase"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres123@localhost:5432/gocommerce?sslmode=disable"
	}

	log.Printf("Connecting to database: %s", dbURL)
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Successfully connected to database")

	repo := infrastructure.NewPostgresUserRepo(db)
	uc := &usecase.UserUsecase{Repo: repo}
	handler := &_interface.UserHandler{Usecase: uc}

	r := gin.Default()
	handler.RegisterRoutes(r)

	log.Println("User Service starting on port 8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
