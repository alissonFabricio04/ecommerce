package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
	"github.com/alissonFabricio04/ecommerce/backend/internal/infra/models"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.CategoryModel{})
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (repo *CategoryRepositoryImpl) Save(category *domain.Category) error {
	categoryModel := &models.CategoryModel{
		Id:   category.Id.String(),
		Name: category.Name,
	}
	tx := repo.db.Create(categoryModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *CategoryRepositoryImpl) FindByID(id *uuid.UUID) (*domain.Category, error) {
	var categoryModel CategoryModel
	tx := repo.db.First(&categoryModel, id.String())
	if tx.Error != nil {
		return nil, tx.Error
	}
	category, err := domain.RestoreCategory(categoryModel.Id, categoryModel.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepositoryImpl) GetAll() ([]*domain.Category, error) {
	var results []*CategoryModel
	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var categories []*domain.Category
	for _, categoryModel := range results {
		category, err := domain.RestoreCategory(categoryModel.Id, categoryModel.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
