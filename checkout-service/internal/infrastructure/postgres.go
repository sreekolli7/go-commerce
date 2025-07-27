package infrastructure

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sreekolli7/checkout-service/internal/domain"
)

type PostgresOrderRepo struct {
	DB *sqlx.DB
}

// I'm creating a new repository instance
// This is how I connect my business logic to the database
func NewPostgresOrderRepo(db *sqlx.DB) *PostgresOrderRepo {
	return &PostgresOrderRepo{DB: db}
}

// I'm saving a new order to the database
// This is where I handle all the database operations for orders
func (r *PostgresOrderRepo) Create(order *domain.Order) error {
	// I'm converting the order items to JSON
	// PostgreSQL can store JSON data, which is really cool
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	// I'm getting the current time for created_at and updated_at
	// This helps me track when orders were created
	now := time.Now()

	// I'm writing the SQL query to insert the order
	// I learned SQL in my database class
	query := `
		INSERT INTO orders (user_id, total, status, items, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, created_at, updated_at
	`

	// I'm executing the query and getting back the generated ID
	// This is how I save the order and get the database-generated ID
	err = r.DB.QueryRowx(
		query,
		order.UserID, // The user who made the order
		order.Total,  // The total price
		order.Status, // The order status (pending, paid, etc.)
		itemsJSON,    // The order items as JSON
		now,          // When the order was created
		now,          // When the order was last updated
	).Scan(&order.ID, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	return nil
}

// I'm getting all orders for a specific user
// This is how I retrieve order history from the database
func (r *PostgresOrderRepo) GetOrdersByUserID(userID int64) ([]domain.Order, error) {
	var orders []domain.Order

	// I'm writing a SQL query to get all orders for a user
	// I'm ordering by created_at DESC to show newest orders first
	query := `
		SELECT id, user_id, total, status, items, created_at, updated_at 
		FROM orders 
		WHERE user_id = $1 
		ORDER BY created_at DESC
	`

	// I'm executing the query and getting all the orders
	err := r.DB.Select(&orders, query, userID)
	return orders, err
}
