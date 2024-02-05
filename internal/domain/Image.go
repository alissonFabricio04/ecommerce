package domain

import (
	"net/url"

	"github.com/google/uuid"
)

type Image struct {
	Id  uuid.UUID
	Uri *url.URL
}

func CreateNewImage(uri string) (*Image, error) {
	uriIsValid, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	return &Image{
		Id:  uuid.New(),
		Uri: uriIsValid,
	}, nil
}

func RestoreImage(id string, uri string) (*Image, error) {
	idIsValid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	uriIsValid, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	return &Image{
		Id:  idIsValid,
		Uri: uriIsValid,
	}, nil
}
