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
