package domain

import "github.com/google/uuid"

type Category struct {
	id   string
	name string
}

func createNewCategoty(name string) *Category {
	return &Category{
		id:   uuid.New().String(),
		name: name,
	}
}
