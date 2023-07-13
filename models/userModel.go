package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int            `gorm:"primaryKey:autoIncrement:1000"`
	FirstName    string         `gorm:"not null"`
	LastName     string         `gorm:"not null"`
	Password     string         `gorm:"not null"`
	Email        sql.NullString `gorm:"unique"`
	MobileNumber string         `gorm:"unique;index;not null"`
	RegisterTime time.Time
	IsVerify     bool
	Category     []Category `gorm:"foreignKey:Id;references:Id"`
	Cards        []Card     `gorm:"foreignKey:Id;references:Id"`
}

func (user *User) TableName() string {
	return "user"
}
