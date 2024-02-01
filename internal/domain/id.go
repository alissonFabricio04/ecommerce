package domain

import "github.com/google/uuid"

type Id struct {
	value string
}

func createNewId() *Id {
	return &Id{
		value: uuid.New().String(),
	}
}
