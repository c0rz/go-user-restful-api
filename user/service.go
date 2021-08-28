package user

type Repository interface {
	getAllUser() ([]User, error)
}
