package _interface

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sreekolli7/go-commerce/user-service/internal/domain"
	"github.com/sreekolli7/go-commerce/user-service/internal/usecase"
)

type UserHandler struct {
	Usecase *usecase.UserUsecase
}

func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.GET("/users/:id", h.GetByID)
}

func (h *UserHandler) Register(c *gin.Context) {
	log.Println("Register endpoint called")

	var req domain.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("JSON binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Registration request for email: %s", req.Email)

	user, err := h.Usecase.Register(&req)
	if err != nil {
		log.Printf("Registration error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("User created successfully with ID: %d", user.ID)

	// Return user data without password
	response := gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}
	c.JSON(http.StatusCreated, response)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Usecase.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetByID(c *gin.Context) {
	// TODO: implement with JWT authentication
	c.JSON(http.StatusOK, gin.H{"message": "Get user by ID - TODO"})
}
