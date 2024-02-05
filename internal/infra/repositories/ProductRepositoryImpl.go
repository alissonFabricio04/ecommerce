package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alissonFabricio04/ecommerce/backend/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func insertBatchImages(repo *ProductRepositoryImpl, productId *uuid.UUID, imgs []*domain.Image) error {
	imgsBatchStrings := make([]string, 0, len(imgs))
	imgsBatchArgs := make([]interface{}, 0, len(imgs)*3)
	for _, img := range imgs {
		imgsBatchStrings = append(imgsBatchStrings, "(?, ?, ?)")
		imgsBatchArgs = append(imgsBatchArgs, img.Id.String())
		imgsBatchArgs = append(imgsBatchArgs, img.Uri.String())
		imgsBatchArgs = append(imgsBatchArgs, productId.String())
	}
	stmt := fmt.Sprintf("INSERT INTO images (id, uri, fk_product) VALUES %s", strings.Join(imgsBatchStrings, ","))
	_, err := repo.db.Exec(stmt, imgsBatchArgs...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) Save(product *domain.Product) error {
	err := insertBatchImages(repo, &product.Id, product.Images)
	if err != nil {
		return err
	}
	_, err = repo.db.Exec("INSERT INTO products (id, name, description, price, fk_category) VALUES (?, ?, ?, ?, ?)", product.Id.String(), product.Name, product.Description, product.Price, product.Category.Id.String())
	if err != nil {
		return err
	}
	return nil
}

func (repo *ProductRepositoryImpl) FindByID(id *uuid.UUID) (*domain.Product, error) {
	// get product
	var product domain.Product
	var categoryId string
	err := repo.db.QueryRow("SELECT id, name, description, price, fk_category FROM products WHERE id = ?", id.String()).Scan(&product.Id, &product.Name, &product.Description, &product.Price, &categoryId)
	if err != nil {
		return nil, err
	}

	// get category of product
	var category domain.Category
	err = repo.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", categoryId).Scan(&category.Id, &category.Name)
	if err != nil {
		return nil, err
	}

	// get images of product
	rows, err := repo.db.Query("SELECT id, uri FROM images WHERE fk_product = ?", id.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var images []*domain.Image
	for rows.Next() {
		var image domain.Image
		if err := rows.Scan(&image.Id, &image.Uri); err != nil {
			return nil, err
		}
		images = append(images, &image)
	}

	// restore category and images
	product.Category = &category
	product.Images = images
	return &product, nil
}

func (repo *ProductRepositoryImpl) FindByCategory(categoryId *uuid.UUID) ([]*domain.Product, error) {
	// get category of product
	var category domain.Category
	err := repo.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", categoryId.String()).Scan(&category.Id, &category.Name)
	if err != nil {
		return nil, err
	}

	// get products
	rows, err := repo.db.Query("SELECT id, name, description, price FROM products WHERE fk_category = ?", categoryId.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		if err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}

		// get images of product
		rows, err := repo.db.Query("SELECT id, uri FROM images WHERE fk_product = ?", product.Id.String())
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var images []*domain.Image
		for rows.Next() {
			var image domain.Image
			if err := rows.Scan(&image.Id, &image.Uri); err != nil {
				return nil, err
			}
			images = append(images, &image)
		}
		product.Category = &category
		product.Images = images
		products = append(products, &product)
	}
	return products, nil
}

func (repo *ProductRepositoryImpl) GetAll() ([]*domain.Product, error) {
	// get products
	rows, err := repo.db.Query("SELECT id, name, description, price, fk_category FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*domain.Product
	for rows.Next() {
		var product domain.Product
		var categoryId string
		if err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &categoryId); err != nil {
			return nil, err
		}

		// get category
		var category domain.Category
		err := repo.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", categoryId).Scan(&category.Id, &category.Name)
		if err != nil {
			return nil, err
		}

		// get images of product
		rows, err := repo.db.Query("SELECT id, uri FROM images WHERE fk_product = ?", product.Id.String())
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		var images []*domain.Image
		for rows.Next() {
			var image domain.Image
			if err := rows.Scan(&image.Id, &image.Uri); err != nil {
				return nil, err
			}
			images = append(images, &image)
		}
		product.Category = &category
		product.Images = images
		products = append(products, &product)
	}
	return products, nil
}
