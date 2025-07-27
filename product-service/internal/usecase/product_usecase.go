package usecase

import (
	"github.com/sreekolli7/go-commerce/product-service/internal/domain"
)

type ProductRepository interface {
	FetchAll() ([]domain.Product, error)
	FetchByID(id int64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int64) error
}

type ProductUsecase struct {
	Repo ProductRepository
}

func (u *ProductUsecase) GetAll() ([]domain.Product, error) {
	return u.Repo.FetchAll()
}

func (u *ProductUsecase) GetByID(id int64) (*domain.Product, error) {
	return u.Repo.FetchByID(id)
}

func (u *ProductUsecase) Create(product *domain.Product) error {
	return u.Repo.Create(product)
}

func (u *ProductUsecase) Update(product *domain.Product) error {
	return u.Repo.Update(product)
}

func (u *ProductUsecase) Delete(id int64) error {
	return u.Repo.Delete(id)
}
