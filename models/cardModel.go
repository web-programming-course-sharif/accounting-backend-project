package models

type Card struct {
	Id            int64
	CardNumber    string
	Bank          Bank   `gorm:"foreignKey:Id;references:Id"`
	AccountNumber string `gorm:"unique"`
	Name          string
	Balance       float64
}

func (card *Card) TableName() string {
	return "card"
}
