package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	GetAllUser() ([]User, error)
	RegisterUser(input UserInput) (User, error)
	CountUser(kolom string, data string) (bool, error)
	UpdateUser(id string, input EditInput) (User, error)
	DeleteUser(id string) error
}

type service struct {
	models Models
}

func NewService(models Models) *service {
	return &service{models}
}

func (s *service) DeleteUser(id string) error {
	delete := s.models.Delete(id)
	if delete != nil {
		return delete
	}
	return nil
}

func (s *service) UpdateUser(id string, input EditInput) (User, error) {
	user, err := s.models.Search("id", id)
	if err != nil {
		return user, err
	}
	if input.Nama != "" {
		user.Nama = input.Nama
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}

		user.Password = string(passwordHash)
	}
	// user.

	updatedUser, err := s.models.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) RegisterUser(input UserInput) (User, error) {
	user := User{}
	user.Nama = input.Nama
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
