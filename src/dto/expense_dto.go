package dto

type Expense struct {
	Id                  uint16               `json:"id"`
	UserId              uint8                `json:"user_id"`
	ReferenceMonth      string               `json:"reference_month"`
	State               string               `json:"state"`
	TotalAmount         int32                `json:"total_amount"`
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
	Id          uint32 `json:"id"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
}

type CreditCardPurchase struct {
	Id                 uint32 `json:"id"`
	Category           string `json:"category"`
	Description        string `json:"description"`
	Amount             int32  `json:"amount"`
	CurrentInstallment uint8  `json:"current_installment"`
	LastInstallment    uint8  `json:"last_installment"`
}

type CreditCardPurchaseRequest struct {
	Id           uint32 `json:"id"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	Amount       int32  `json:"amount"`
	Installments uint8  `json:"installments"`
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
