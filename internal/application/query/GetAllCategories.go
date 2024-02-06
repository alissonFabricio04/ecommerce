package query

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/application/repositories"
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
)

type GetAllCategoriesQuery struct {
	categoryRepository repositories.CategoryRepository
}

func InstaceNewGetAllCategoriesQuery(categoryRepository repositories.CategoryRepository) *GetAllCategoriesQuery {
	return &GetAllCategoriesQuery{
		categoryRepository: categoryRepository,
	}
}

func (query *GetAllCategoriesQuery) Handle() ([]*domain.Category, error) {
	categories, err := query.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return categories, nil
}
