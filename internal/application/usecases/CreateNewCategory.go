package usecases

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/application/repositories"
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
)

type CreateNewCategoryUseCase struct {
	categoryRepository repositories.CategoryRepository
}

func InstaceNewCreateNewCategoryUseCase(categoryRepository repositories.CategoryRepository) *CreateNewCategoryUseCase {
	return &CreateNewCategoryUseCase{
		categoryRepository: categoryRepository,
	}
}

func (useCase *CreateNewCategoryUseCase) Handle(name string) (*domain.Id, error) {
	category, err := domain.CreateNewCategoty(name)
	if err != nil {
		return nil, err
	}
	if err = useCase.categoryRepository.Save(category); err != nil {
		return nil, err
	}
	return category.Id, nil
}
