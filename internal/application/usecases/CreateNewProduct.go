package usecases

import (
	"errors"

	"github.com/alissonFabricio04/ecommerce/backend/internal/application/repositories"
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
)

type CreateNewProductUseCase struct {
	categoryRepository repositories.CategoryRepository
	productRepository  repositories.ProductRepository
}

func InstaceNewCreateNewProductUseCase(
	categoryRepository repositories.CategoryRepository,
	productRepository repositories.ProductRepository,
) *CreateNewProductUseCase {
	return &CreateNewProductUseCase{
		categoryRepository: categoryRepository,
		productRepository:  productRepository,
	}
}

func (useCase *CreateNewProductUseCase) Handle(name string, description string, price float64, categoryId string, imgs []string) (*domain.Id, error) {
	categId, err := domain.InstanceNewId(categoryId)
	if err != nil {
		return nil, err
	}
	category, err := useCase.categoryRepository.FindById(categId)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("the category selected does not exists")
	}
	var images []*domain.Image
	for i := 0; i < len(imgs); i++ {
		image, err := domain.CreateNewImage(imgs[i])
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	product, err := domain.CreateNewProduct(name, description, price, category, images)
	if err != nil {
		return nil, err
	}
	if err = useCase.productRepository.Save(product); err != nil {
		return nil, err
	}
	return product.Id, nil
}
