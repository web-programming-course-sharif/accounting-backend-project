package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type CardRepository interface {
	CreateCard(card models.Card) (models.Card, error)
	GetAllCards(userId int) []models.Card
}

func RepositoryCard(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) CreateCard(card models.Card) (models.Card, error) {
	err := r.db.Create(&card).Error
	return card, err
}
func (r *repository) GetAllCards(userId int) []models.Card {
	var cards []models.Card
	r.db.Where("user_id = ?", userId).Find(&cards)
	// SELECT * FROM cards WHERE user_id = id OR user_id=nil;
	return cards
}
