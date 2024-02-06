package domain

type Category struct {
	Id   *Id
	Name *Name
}

func CreateNewCategoty(name string) (*Category, error) {
	nameIsValid, err := InstanceNewName(name)
	if err != nil {
		return nil, err
	}
	return &Category{
		Id:   CreateNewId(),
		Name: nameIsValid,
	}, nil
}

func RestoreCategory(id *Id, name string) (*Category, error) {
	nameIsValid, err := InstanceNewName(name)
	if err != nil {
		return nil, err
	}
	return &Category{
		Id:   id,
		Name: nameIsValid,
	}, nil
}
