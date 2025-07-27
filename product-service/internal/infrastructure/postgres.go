package infrastructure

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sreekolli7/go-commerce/product-service/internal/domain"
	"github.com/sreekolli7/go-commerce/product-service/internal/usecase"
)

type PostgresProductRepo struct {
	DB *sqlx.DB
}

func NewPostgresProductRepo(db *sqlx.DB) usecase.ProductRepository {
	return &PostgresProductRepo{DB: db}
}

func (r *PostgresProductRepo) FetchAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.DB.Select(&products, "SELECT * FROM products ORDER BY id")
	return products, err
}

func (r *PostgresProductRepo) FetchByID(id int64) (*domain.Product, error) {
	var product domain.Product
	err := r.DB.Get(&product, "SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *PostgresProductRepo) Create(product *domain.Product) error {
	return r.DB.QueryRowx(
		"INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) RETURNING id",
		product.Name, product.Description, product.Price, product.Stock,
	).Scan(&product.ID)
}

func (r *PostgresProductRepo) Update(product *domain.Product) error {
	_, err := r.DB.Exec(
		"UPDATE products SET name = $1, description = $2, price = $3, stock = $4 WHERE id = $5",
		product.Name, product.Description, product.Price, product.Stock, product.ID,
	)
	return err
}

func (r *PostgresProductRepo) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
