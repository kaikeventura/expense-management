package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        uint8     `gorm:"primaryKey; autoIncrement"`
	Username  string    `gorm:"not null, unique"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"null"`
}
