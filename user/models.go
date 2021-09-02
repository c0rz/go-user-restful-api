package user

import "gorm.io/gorm"

type Models interface {
	GetAll() ([]User, error)
	Insert(user User) (User, error)
	Search(kolom string, data string) (User, error)
	Update(user User) (User, error)
	Delete(id string) error
}

type models struct {
	db *gorm.DB
}

func ConnectDB(db *gorm.DB) *models {
	db.AutoMigrate(&User{})
	return &models{db}
}

func (r *models) Delete(id string) error {
	error := r.db.Delete(&User{}, id).Error
	if error != nil {
		return error
	}
	return nil
}

func (r *models) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
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
