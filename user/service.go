package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	GetAllUser() ([]User, error)
	RegisterUser(input RegisterUserInput) (User, error)
	CountUser(kolom string, data string) (bool, error)
}

type service struct {
	models Models
}

func NewService(models Models) *service {
	return &service{models}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Nama = input.Name
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.models.Insert(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GetAllUser() ([]User, error) {
	users, err := s.models.GetAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) CountUser(kolom string, data string) (bool, error) {
	users, err := s.models.Search(kolom, data)
	if err != nil {
		return false, err
	}

	if users.ID == 0 {
		return true, nil
	}

	return false, nil
}
