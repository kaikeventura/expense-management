package entity

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Id                  uint16               `gorm:"primaryKey; autoIncrement"`
	UserId              uint8                `gorm:"not null"`
	Reference           time.Time            `gorm:"not null"`
	State               string               `gorm:"not null"`
	TotalAmount         uint32               `gorm:"not null"`
	FixedExpenses       []FixedExpense       `gorm:"null"`
	Purchases           []Purchase           `gorm:"null"`
	CreditCardPurchases []CreditCardPurchase `gorm:"null"`
}

type FixedExpense struct {
	gorm.Model
	Id          uint32 `gorm:"primaryKey; autoIncrement"`
	ExpenseId   uint16 `gorm:"null"`
	Category    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Amount      int32  `gorm:"not null"`
}

type Purchase struct {
	gorm.Model
	Id          uint32 `gorm:"primaryKey; autoIncrement"`
	ExpenseId   uint16 `gorm:"null"`
	Category    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Amount      int32  `gorm:"not null"`
}

type CreditCardPurchase struct {
	gorm.Model
	Id                 uint32 `gorm:"primaryKey; autoIncrement"`
	ExpenseId          uint16 `gorm:"null"`
	Category           string `gorm:"not null"`
	Description        string `gorm:"not null"`
	Amount             int32  `gorm:"not null"`
	CurrentInstallment uint8  `gorm:"not null"`
	LastInstallment    uint8  `gorm:"not null"`
}
