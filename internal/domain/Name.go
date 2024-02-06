package domain

import (
	"errors"
)

type Name struct {
	Value string
}

func InstanceNewName(name string) (*Name, error) {
	if len(name) <= 1 {
		return nil, errors.New("the name is much short")
	}
	if len(name) > 120 {
		return nil, errors.New("the name is much long")
	}
	return &Name{
		Value: name,
	}, nil
}
