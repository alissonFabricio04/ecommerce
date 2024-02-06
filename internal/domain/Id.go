package domain

import (
	"github.com/google/uuid"
)

type Id struct {
	Value uuid.UUID
}

func CreateNewId() *Id {
	return &Id{
		Value: uuid.New(),
	}
}

func InstanceNewId(id string) (*Id, error) {
	idIsValid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &Id{
		Value: idIsValid,
	}, nil
}

func (id *Id) ToString() string {
	return id.Value.String()
}
