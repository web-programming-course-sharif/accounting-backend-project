package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type CardRepository interface {
	createCard(id int) models.User
}

func RepositoryCard(db *gorm.DB) *repository {
	return &repository{db}
}
