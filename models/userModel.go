package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int            `json:"id"  gorm:"primaryKey:autoIncrement:1000"`
	FirstName    string         `json:"firstName" gorm:"not null"`
	LastName     string         `json:"lastName" gorm:"not null"`
	Country      string         `json:"country"`
	Address      string         `json:"address"`
	PhotoURL     string         `json:"photoURL"`
	About        string         `json:"about"`
	IsPublic     bool           `json:"isPublic" gorm:"default:true"`
	State        string         `json:"state"`
	City         string         `json:"city"`
	Password     string         `json:"password" gorm:"not null"`
	Email        sql.NullString `json:"email" gorm:"unique"`
	PhoneNumber  string         `json:"phoneNumber" gorm:"unique;index;not null"`
	RegisterTime time.Time      `json:"registerTime"`
	IsVerify     bool           `json:"isVerify"`
	Category     []Category     `json:"category" gorm:"foreignKey:Id;references:Id"`
	Cards        []Card         `json:"cards" gorm:"foreignKey:Id;references:Id"`
}

func (user *User) TableName() string {
	return "user"
}
