package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&CategoryModel{})
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (repo *CategoryRepositoryImpl) Save(category *domain.Category) error {
	categoryModel := &CategoryModel{
		Id:   category.Id.ToString(),
		Name: category.Name.Value,
	}
	tx := repo.db.Create(&categoryModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *CategoryRepositoryImpl) FindById(id *domain.Id) (*domain.Category, error) {
	var categoryModel CategoryModel
	tx := repo.db.First(&categoryModel, id.ToString())
	if tx.Error != nil {
		return nil, tx.Error
	}
	categoryId, err := domain.InstanceNewId(categoryModel.Id)
	if err != nil {
		return nil, err
	}
	category, err := domain.RestoreCategory(categoryId, categoryModel.Name)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepositoryImpl) GetAll() ([]*domain.Category, error) {
	var categoriesModel []*CategoryModel
	tx := repo.db.Find(&categoriesModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var categories []*domain.Category
	for _, categoryModel := range categoriesModel {
		categoryId, err := domain.InstanceNewId(categoryModel.Id)
		if err != nil {
			return nil, err
		}
		category, err := domain.RestoreCategory(categoryId, categoryModel.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
