package _interface

import (
	"net/http"
	"strconv"

	"github.com/sreekolli7/checkout-service/internal/domain"
	"github.com/sreekolli7/checkout-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler struct {
	Usecase *usecase.CheckoutUsecase
}

func (h *CheckoutHandler) RegisterRoutes(r *gin.Engine) {
	// I'm creating a group for checkout routes
	// This helps me organize my API endpoints
	auth := r.Group("/checkout")
	auth.Use(JWTAuthMiddleware())
	auth.POST("/", h.Checkout)
	auth.GET("/orders/:user_id", h.GetOrdersByUserID)
}

// I'm creating a placeholder for JWT authentication
// In a real app, I would implement proper JWT validation here
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: I need to implement JWT validation here
		// For now, I'm just letting all requests through
		c.Next()
	}
}

func (h *CheckoutHandler) GetOrdersByUserID(c *gin.Context) {
	// I'm getting the user ID from the URL parameter
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	// I'm calling my business logic to get the orders
	orders, err := h.Usecase.GetOrdersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// I'm sending back the orders as JSON
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (h *CheckoutHandler) Checkout(c *gin.Context) {
	// I'm reading the JSON request from the client
	// This is how I get the order data from the frontend
	var req domain.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// I'm calling my business logic to process the checkout
	// This is where all the order processing happens
	order, err := h.Usecase.Checkout(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// I'm sending back the created order as JSON
	// The frontend will use this to show the order confirmation
	c.JSON(http.StatusCreated, order)
}
