package models

import (
	"time"
)

type User struct {
	Id            int        `json:"id"  gorm:"primaryKey:autoIncrement:1000"`
	FirstName     string     `json:"firstName" gorm:"not null"`
	LastName      string     `json:"lastName" gorm:"not null"`
	Country       string     `json:"country" gorm:"default:null"`
	Address       string     `json:"address" gorm:"default:null"`
	PhotoURL      string     `json:"photoURL" gorm:"default:null"`
	About         string     `json:"about" gorm:"default:null"`
	IsPublic      bool       `json:"isPublic" gorm:"default:true"`
	State         string     `json:"state" gorm:"default:null"`
	City          string     `json:"city" gorm:"default:null"`
	Role          string     `json:"role" gorm:"default:null"`
	ZipCode       string     `json:"zipCode" gorm:"default:null"`
	FacebookLink  string     `json:"facebookLink" gorm:"default:null"`
	InstagramLink string     `json:"instagramLink" gorm:"default:null"`
	LinkedinLink  string     `json:"linkedinLink" gorm:"default:null"`
	TwitterLink   string     `json:"twitterLink" gorm:"default:null"`
	Password      string     `json:"password" gorm:"not null"`
	Email         string     `json:"email" gorm:"unique;default:null"`
	PhoneNumber   string     `json:"phoneNumber" gorm:"unique;index;not null"`
	RegisterTime  time.Time  `json:"registerTime"`
	IsVerify      bool       `json:"isVerify"`
	Category      []Category `json:"category" gorm:"foreignKey:Id;references:Id"`
	Cards         []Card     `json:"cards" gorm:"foreignKey:Id;references:Id"`
}

func (user *User) TableName() string {
	return "user"
}
