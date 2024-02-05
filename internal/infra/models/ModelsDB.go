package repositories

import "gorm.io/gorm"

type CategoryModel struct {
	gorm.Model
	Id   string
	Name string
}

type Product struct {
	gorm.Model
	Id          string
	Name        string
	Description string
	Price       float64
	Fk_category string
}

type Image struct {
	gorm.Model
	Id         string
	Uri        string
	Fk_product string
}
