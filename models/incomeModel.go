package models

type Income struct {
	Id     uint    `json:"id" gorm:"primaryKey:autoIncrement"`
	UserId int64   `json:"userId"`
	User   User    `json:"user" gorm:"foreignKey:UserId;references:Id"`
	Price  float64 `json:"price"`
	Date   string  `json:"date"`
}
