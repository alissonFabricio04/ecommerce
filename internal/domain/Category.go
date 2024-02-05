package domain

import "github.com/google/uuid"

type Category struct {
	Id   uuid.UUID
	Name string
}

func CreateNewCategoty(name string) *Category {
	return &Category{
		Id:   uuid.New(),
		Name: name,
	}
}

func RestoreCategory(id string, name string) (*Category, error) {
	idIsValid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &Category{
		Id:   idIsValid,
		Name: name,
	}, nil
}
