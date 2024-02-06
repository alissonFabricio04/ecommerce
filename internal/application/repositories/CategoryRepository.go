package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
)

type CategoryRepository interface {
	Save(category *domain.Category) error
	FindById(id *domain.Id) (*domain.Category, error)
	GetAll() ([]*domain.Category, error)
}
