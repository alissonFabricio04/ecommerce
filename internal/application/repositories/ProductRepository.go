package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
	"github.com/google/uuid"
)

type ProductRepository interface {
	Save(product *domain.Product) error
	FindById(id *uuid.UUID) (*domain.Product, error)
	FindByCategory(categoryId *uuid.UUID) ([]*domain.Product, error)
	GetAll() ([]*domain.Product, error)
}
