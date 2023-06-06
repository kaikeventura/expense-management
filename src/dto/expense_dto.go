package dto

type Expense struct {
	Id                  uint16               `json:"id"`
	UserId              uint8                `json:"user_id"`
	ReferenceMonth      string               `json:"reference_month"`
	State               string               `json:"state"`
	TotalAmount         uint32               `json:"total_amount"`
	FixedExpenses       []FixedExpense       `json:"fixed_expenses"`
	Purchases           []Purchase           `json:"purchases"`
	CreditCardPurchases []CreditCardPurchase `json:"credit_card_purchases"`
}

type FixedExpense struct {
	Id          uint32 `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
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

func (s Category) CategoryToString() string {
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

type State int

const (
	CLOSED State = iota
	CURRENT
	FUTURE
)

func (s State) StateToString() string {
	switch s {
	case CLOSED:
		return "CLOSED"
	case CURRENT:
		return "CURRENT"
	case FUTURE:
		return "FUTURE"
	default:
		return "INVALID"
	}
}

const (
	YYYYMM = "2006-01"
)
