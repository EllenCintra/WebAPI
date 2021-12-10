package entity

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	CreatedAt   time.Time `gorm:"<-:create"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"description"`
	MediumPrice float32   `gorm:"medium_price"`
	Author      string    `gorm:"author"`
	ImageURL    string    `gorm:"img_url"`
	UserID      int       `gorm:"user_id"`
}
