package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint8     `gorm:"primaryKey; autoIncrement"`
	Username string    `gorm:"not null; unique"`
	Expenses []Expense `gorm:"null"`
}
