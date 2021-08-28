package user

import "gorm.io/gorm"

type Models interface {
	GetAll() ([]User, error)
}

type models struct {
	db *gorm.DB
}

func ConnectDB(db *gorm.DB) *models {
	db.AutoMigrate(&User{})
	return &models{db}
}

func (r *models) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
