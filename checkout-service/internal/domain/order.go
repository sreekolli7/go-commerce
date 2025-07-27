package domain

import "time"

// I'm defining what an order looks like in my system
// This is called a struct - it's like a template for order data
type Order struct {
	ID        int64       `json:"id" db:"id"`                 // The unique order ID from the database
	UserID    int64       `json:"user_id" db:"user_id"`       // Which user made this order
	Items     []OrderItem `json:"items"`                      // The items in the order
	Total     float64     `json:"total" db:"total"`           // The total price of the order
	Status    string      `json:"status" db:"status"`         // The order status (pending, paid, shipped, etc.)
	CreatedAt time.Time   `json:"created_at" db:"created_at"` // When the order was created
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"` // When the order was last updated
}

// I'm defining what an order item looks like
// This represents a single product in an order
type OrderItem struct {
	ID        int64     `json:"id" db:"id"`                 // The unique item ID
	OrderID   int64     `json:"order_id" db:"order_id"`     // Which order this item belongs to
	ProductID int64     `json:"product_id" db:"product_id"` // Which product this is
	Quantity  int       `json:"quantity" db:"quantity"`     // How many of this product
	Price     float64   `json:"price" db:"price"`           // The price per unit
	CreatedAt time.Time `json:"created_at" db:"created_at"` // When this item was added
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // When this item was last updated
}

// I'm defining what a checkout request looks like
// This is what the frontend sends when someone wants to create an order
type CheckoutRequest struct {
	UserID        int64       `json:"user_id"`        // Which user is making the order
	Items         []OrderItem `json:"items"`          // What items they want to buy
	PaymentMethod string      `json:"payment_method"` // How they want to pay (credit card, etc.)
}

type OrderRepository interface {
	Save(order *Order) error
	GetByUserID(userID int64) ([]Order, error)
}
