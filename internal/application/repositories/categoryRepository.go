package repositories

type CategoryRepository interface {
	findById(id *Id) string
}
