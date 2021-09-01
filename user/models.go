package user

import "gorm.io/gorm"

type Models interface {
	GetAll() ([]User, error)
	Insert(user User) (User, error)
	Search(kolom string, data string) (User, error)
}

type models struct {
	db *gorm.DB
}

func ConnectDB(db *gorm.DB) *models {
	db.AutoMigrate(&User{})
	return &models{db}
}

func (r *models) Search(kolom string, data string) (User, error) {
	var users User
	err := r.db.Where(kolom+" = ?", data).Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *models) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *models) Insert(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
