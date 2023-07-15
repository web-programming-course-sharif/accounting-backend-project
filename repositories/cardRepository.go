package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type CardRepository interface {
	CreateCard(card models.Card) (models.Card, error)
}

func RepositoryCard(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreateCard(card models.Card) (models.Card, error) {
	err := r.db.Create(&card).Error
	return card, err
}
