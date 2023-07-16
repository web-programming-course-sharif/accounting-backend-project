package cardDto

type CreateCardRequest struct {
	CardNumber string  `json:"card_number" validation:"required"`
	BankId     int     `json:"bank_id" validation:"required"`
	Balance    float64 `json:"balance" validation:"required"`
	Name       string  `json:"name" validation:"required"`
}
type DeleteCardRequest struct {
	CardId int `json:"card_id" validation:"required"`
}
type EditCardRequest struct {
	CardId     int     `json:"card_id" validation:"required"`
	CardNumber string  `json:"card_number" validation:"required"`
	BankId     int     `json:"bank_id" validation:"required"`
	Balance    float64 `json:"balance" validation:"required"`
	Name       string  `json:"name" validation:"required"`
}
