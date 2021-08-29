package user

type Service interface {
	GetAllUser() ([]User, error)
}

type service struct {
	models Models
}

func NewService(models Models) *service {
	return &service{models}
}

func (s *service) GetAllUser() ([]User, error) {
	users, err := s.models.GetAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
