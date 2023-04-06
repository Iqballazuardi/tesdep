package database

import (
	"deploy/features/user/repository"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(repository.User{})
}
