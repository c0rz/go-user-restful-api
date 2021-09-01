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

type UserInput struct {
	ID       int
	Nama     string `json:"nama" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type EditInput struct {
	ID       int
	Nama     string `json:"name"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}
