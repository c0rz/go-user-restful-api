package user

type Service interface {
	getAllUser() ([]User, error)
}

type service struct {
	models Models
}

func NewService(models Models) *service {
	return &service{models}
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.models.GetAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
