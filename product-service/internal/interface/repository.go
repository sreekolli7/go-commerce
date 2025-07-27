package _interface

import "github.com/sreekolli7/go-commerce/product-service/internal/domain"

type ProductRepository interface {
	FetchAll() ([]domain.Product, error)
	FetchByID(id int64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int64) error
}
