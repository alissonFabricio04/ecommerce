package domain

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID
	Name        string
	Description string
	Price       float64
	Category    *Category
	Images      []*Image
}

func CreateNewProduct(name string, description string, price float64, category *Category, imgs []*Image) (*Product, error) {
	// var images []*Image
	// for i := 0; i < len(imgs); i++ {
	// 	image, err := CreateNewImage(imgs[i])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	images = append(images, image)
	// }
	return &Product{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Images:      imgs,
	}, nil
}

func RestoreProduct(id string, name string, description string, price float64, category *Category, imgs []*Image) (*Product, error) {
	idIsValid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &Product{
		Id:          idIsValid,
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		Images:      imgs,
	}, nil
}
