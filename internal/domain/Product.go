package domain

import (
	"errors"
)

type Product struct {
	Id          *Id
	Name        *Name
	Description string
	Price       float64
	Category    *Category
	Images      []*Image
}

func CreateNewProduct(name string, description string, price float64, category *Category, imgs []*Image) (*Product, error) {
	nameIsValid, err := InstanceNewName(name)
	if err != nil {
		return nil, err
	}
	if len(description) <= 0 {
		return nil, errors.New("the description is much short")
	}
	if len(description) > 120 {
		return nil, errors.New("the description is much long")
	}
	if price <= 0 {
		return nil, errors.New("the price is valid")
	}
	return &Product{
		Id:          CreateNewId(),
		Name:        nameIsValid,
		Description: description,
		Price:       price,
		Category:    category,
		Images:      imgs,
	}, nil
}

func RestoreProduct(id *Id, name string, description string, price float64, category *Category, imgs []*Image) (*Product, error) {
	nameIsValid, err := InstanceNewName(name)
	if err != nil {
		return nil, err
	}
	return &Product{
		Id:          id,
		Name:        nameIsValid,
		Description: description,
		Price:       price,
		Category:    category,
		Images:      imgs,
	}, nil
}
