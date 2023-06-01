package dto

import "time"

type Expense struct {
	Id                  uint16               `json:"id"`
	UserId              uint8                `json:"user_id"`
	Reference           time.Time            `json:"reference"`
	State               string               `json:"state"`
	TotalAmount         uint32               `json:"total_amount"`
	FixedExpenses       []FixedExpense       `json:"fixed_expenses"`
	Purchases           []Purchase           `json:"purchases"`
	CreditCardPurchases []CreditCardPurchase `json:"credit_card_purchases"`
}

type FixedExpense struct {
	Id          uint32   `json:"id"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
	Amount      int32    `json:"amount"`
}

type Purchase struct {
	Id          uint32   `json:"id"`
	Category    Category `json:"category"`
	Description string   `json:"description"`
	Amount      int32    `json:"amount"`
}

type CreditCardPurchase struct {
	Id                 uint32   `json:"id"`
	Category           Category `json:"category"`
	Description        string   `json:"description"`
	Amount             int32    `json:"amount"`
	CurrentInstallment uint8    `json:"current_installment"`
	LastInstallment    uint8    `json:"last_installment"`
}

type Category int

const (
	Unknown Category = iota
	Pending
	InProgress
	Completed
)

func (s Category) String() string {
	switch s {
	case Unknown:
		return "Unknown"
	case Pending:
		return "Pending"
	case InProgress:
		return "In Progress"
	case Completed:
		return "Completed"
	default:
		return "Invalid"
	}
}
