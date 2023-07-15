package card

type createCard struct {
	CardNumber string `json:"card_number" validation:"required"`
	BankName   string `json:"bank_name" validation:"required"`
	Balance    string `json:"balance" validation:"required"`
}
