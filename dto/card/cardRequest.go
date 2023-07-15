package cardDto

type CreateCardRequest struct {
	CardNumber    string  `json:"card_number" validation:"required"`
	BankId        int     `json:"bank_id" validation:"required"`
	Balance       float64 `json:"balance" validation:"required"`
	AccountNumber string  `json:"account_number"`
	Name          string  `json:"name" validation:"required"`
}
