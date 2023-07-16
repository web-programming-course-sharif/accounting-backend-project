package models

type Card struct {
	Id            int64   `json:"id"`
	UserId        int64   `json:"userId"`
	User          User    `json:"user" gorm:"foreignKey:UserId;references:Id"`
	CardNumber    string  `json:"cardNumber" gorm:"unique"`
	BankId        int     `json:"bankId"`
	Bank          Bank    `json:"bank" gorm:"foreignKey:BankId;references:Id"`
	AccountNumber string  `json:"accountNumber"`
	Name          string  `json:"name"`
	Balance       float64 `json:"balance"`
}

func (card *Card) TableName() string {
	return "card"
}
