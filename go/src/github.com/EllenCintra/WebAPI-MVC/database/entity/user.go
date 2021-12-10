package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	CreatedAt time.Time `gorm:"<-:create"`
	Name      string    `gorm:"name"`
	Email     string    `gorm:"email"`
	Password  string    `gorm:"password"`
	Book      []Book
}
