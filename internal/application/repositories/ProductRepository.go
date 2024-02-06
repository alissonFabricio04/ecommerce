package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
)

type ProductRepository interface {
	Save(product *domain.Product) error
	FindById(id *domain.Id) (*domain.Product, error)
	FindByCategory(categoryId *domain.Id) ([]*domain.Product, error)
	GetAll() ([]*domain.Product, error)
}
