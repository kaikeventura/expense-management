package entity

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Id                  uint16               `gorm:"primaryKey; autoIncrement"`
	UserId              uint8                `gorm:"not null"`
	ReferenceMonth      string               `gorm:"not null; size:7"`
	SequenceNumber      uint16               `gorm:"not null"`
	State               string               `gorm:"not null"`
	TotalAmount         int32                `gorm:"not null"`
	FixedExpenses       []FixedExpense       `gorm:"null, foreignKey:ExpenseId"`
	Purchases           []Purchase           `gorm:"null, foreignKey:ExpenseId"`
	CreditCardPurchases []CreditCardPurchase `gorm:"null, foreignKey:ExpenseId"`
}

type FixedExpense struct {
	gorm.Model
	Id          uint32 `gorm:"primaryKey; autoIncrement"`
	ExpenseId   uint16 `gorm:"not null"`
	Category    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Amount      int32  `gorm:"not null"`
}

type Purchase struct {
	gorm.Model
	Id          uint32 `gorm:"primaryKey; autoIncrement"`
	ExpenseId   uint16 `gorm:"not null"`
	Category    string `gorm:"not null"`
	Description string `gorm:"not null"`
	Amount      int32  `gorm:"not null"`
}

type CreditCardPurchase struct {
	gorm.Model
	Id                 uint32 `gorm:"primaryKey; autoIncrement"`
	ExpenseId          uint16 `gorm:"not null"`
	Category           string `gorm:"not null"`
	Description        string `gorm:"not null"`
	Amount             int32  `gorm:"not null"`
	CurrentInstallment uint8  `gorm:"not null"`
	LastInstallment    uint8  `gorm:"not null"`
}

func PlusExpenseTotalAmount(expense *Expense, amount int32) {
	expense.TotalAmount += amount
}
