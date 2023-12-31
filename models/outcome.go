package models

type Outcome struct {
	Id         uint     `json:"id" gorm:"primaryKey:autoIncrement"`
	UserId     int64    `json:"userId"`
	User       User     `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Price      float64  `json:"price"`
	Type       int      `json:"type"`
	Date       string   `json:"date"`
	CardId     int      `json:"cardId"`
	Card       Card     `json:"card" gorm:"foreignKey:CardId;references:Id"`
	CategoryId int      `json:"CategoryId"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId;references:Id"`
}

func (outcome *Outcome) TableName() string {
	return "outcome"
}
