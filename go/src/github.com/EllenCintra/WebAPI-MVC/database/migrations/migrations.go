package migrations

import (
	"github.com/hyperyuri/webapi-with-go/database/entity"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(entity.Book{})
	db.AutoMigrate(entity.User{})
}
