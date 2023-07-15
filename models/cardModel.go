package models

type Card struct {
	Id            int64
	UserId        int64 `json:"userId"`
	User          User  `gorm:"foreignKey:UserId;references:Id"`
	CardNumber    string
	BankId        int
	Bank          Bank   `gorm:"foreignKey:BankId;references:Id"`
	AccountNumber string `gorm:"unique"`
	Name          string
	Balance       float64
}

func (card *Card) TableName() string {
	return "card"
}
