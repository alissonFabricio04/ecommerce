package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	Save(category *domain.Category) error
	FindById(id *uuid.UUID) (*domain.Category, error)
	GetAll() ([]*domain.Category, error)
}
