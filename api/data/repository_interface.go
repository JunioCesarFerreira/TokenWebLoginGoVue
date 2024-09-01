package data

type Repository interface {
	CheckUser(id string, hash string) (bool, int, error)
}
