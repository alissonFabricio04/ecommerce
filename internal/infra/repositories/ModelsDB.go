package repositories

import "gorm.io/gorm"

type CategoryModel struct {
	gorm.Model
	Id   string `gorm:"primaryKey"`
	Name string
}

func (CategoryModel) TableName() string {
	return "categories"
}

type ProductModel struct {
	gorm.Model
	Id          string `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Fk_category string `gorm:"foreignKey:categories"`
}

func (ProductModel) TableName() string {
	return "products"
}

type ImageModel struct {
	gorm.Model
	Id         string `gorm:"primaryKey"`
	Uri        string
	Fk_product string `gorm:"foreignKey:products"`
}

func (ImageModel) TableName() string {
	return "images_products"
}
