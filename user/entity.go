package user

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
