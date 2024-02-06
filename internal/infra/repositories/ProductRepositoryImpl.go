package repositories

import (
	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&ProductModel{})
	return &ProductRepositoryImpl{
		db: db,
	}
}

func insertBatchImages(repo *ProductRepositoryImpl, productId *domain.Id, imgs []*domain.Image) error {
	var imagesModel []*ImageModel
	for i := 0; i < len(imgs); i++ {
		imageModel := &ImageModel{
			Id:         imgs[i].Id.String(),
			Uri:        imgs[i].Uri.String(),
			Fk_product: productId.ToString(),
		}
		imagesModel = append(imagesModel, imageModel)
	}
	tx := repo.db.Create(&imagesModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ProductRepositoryImpl) Save(product *domain.Product) error {
	err := insertBatchImages(repo, product.Id, product.Images)
	if err != nil {
		return err
	}
	productModel := &ProductModel{
		Id:          product.Id.ToString(),
		Name:        product.Name.Value,
		Description: product.Description,
		Price:       product.Price,
		Fk_category: product.Category.Id.ToString(),
	}
	tx := repo.db.Create(&productModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *ProductRepositoryImpl) FindById(id *domain.Id) (*domain.Product, error) {
	// get product
	var productModel ProductModel
	tx := repo.db.First(&productModel, id.ToString())
	if tx.Error != nil {
		return nil, tx.Error
	}

	// get category
	var categoryModel CategoryModel
	tx = repo.db.First(&categoryModel, productModel.Fk_category)
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

	// get images
	var imagesModel []*ImageModel
	tx = repo.db.Find(&imagesModel, "fk_product = ?", productModel.Id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var images []*domain.Image
	for _, imageModel := range imagesModel {
		image, err := domain.RestoreImage(imageModel.Id, imageModel.Uri)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	// restore product
	productId, err := domain.InstanceNewId(productModel.Id)
	if err != nil {
		return nil, err
	}
	product, err := domain.RestoreProduct(
		productId,
		productModel.Name,
		productModel.Description,
		productModel.Price,
		category,
		images,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (repo *ProductRepositoryImpl) FindByCategory(categoryId *domain.Id) ([]*domain.Product, error) {
	// get category
	var categoryModel CategoryModel
	tx := repo.db.First(&categoryModel, categoryId.ToString())
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
	// get product
	var productsModel []*ProductModel
	tx = repo.db.First(&productsModel, "fk_category = ?", categoryId.ToString())
	if tx.Error != nil {
		return nil, tx.Error
	}
	var products []*domain.Product
	for _, productModel := range productsModel {
		// get images
		var imagesModel []*ImageModel
		tx = repo.db.Find(&imagesModel, "fk_product = ?", productModel.Id)
		if tx.Error != nil {
			return nil, tx.Error
		}
		var images []*domain.Image
		for _, imageModel := range imagesModel {
			image, err := domain.RestoreImage(imageModel.Id, imageModel.Uri)
			if err != nil {
				return nil, err
			}
			images = append(images, image)
		}

		// restore product
		productId, err := domain.InstanceNewId(productModel.Id)
		if err != nil {
			return nil, err
		}
		product, err := domain.RestoreProduct(
			productId,
			productModel.Name,
			productModel.Description,
			productModel.Price,
			category,
			images,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepositoryImpl) GetAll() ([]*domain.Product, error) {
	var productsModel []*ProductModel
	tx := repo.db.First(&productsModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var products []*domain.Product
	for _, productModel := range productsModel {
		// get category
		var categoryModel CategoryModel
		tx = repo.db.First(&categoryModel, productModel.Fk_category)
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

		// get images
		var imagesModel []*ImageModel
		tx = repo.db.Find(&imagesModel, "fk_product = ?", productModel.Id)
		if tx.Error != nil {
			return nil, tx.Error
		}
		var images []*domain.Image
		for _, imageModel := range imagesModel {
			image, err := domain.RestoreImage(imageModel.Id, imageModel.Uri)
			if err != nil {
				return nil, err
			}
			images = append(images, image)
		}
		// restore product
		productId, err := domain.InstanceNewId(productModel.Id)
		if err != nil {
			return nil, err
		}
		product, err := domain.RestoreProduct(
			productId,
			productModel.Name,
			productModel.Description,
			productModel.Price,
			category,
			images,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
