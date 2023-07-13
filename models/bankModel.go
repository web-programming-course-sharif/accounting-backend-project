package models

type Bank struct {
	Id   uint `gorm:"primaryKey:autoIncrement"`
	Name string
	Icon string
}

func (bank *Bank) TableName() string {
	return "bank"
}
