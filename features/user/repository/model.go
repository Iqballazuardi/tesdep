package repository

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string
	Nama      string
	HP        string `gorm:"primaryKey;type:varchar(13);"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
