package domain

import "github.com/google/uuid"

type Product struct {
	id          string
	name        string
	description string
	price       float64
	category    *Category
	imagesUri   []string
}

func createNewProduct(name string, description string, price float64, category *Category, imagesUri []string) *Product {
	return &Product{
		id:          uuid.New().String(),
		name:        name,
		description: description,
		price:       price,
		category:    category,
		imagesUri:   imagesUri,
	}
}
