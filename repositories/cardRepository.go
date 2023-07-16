package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type CardRepository interface {
	CreateCard(card models.Card) (models.Card, error)
	GetAllCards(userId int) []models.Card
	DeleteCard(cardId int) error
	GetCardById(cardId int) models.Card
	EditCard(id int, cardModel models.Card) (models.Card, error)
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
	r.db.Where("user_id = ?", userId).Preload("Bank").Find(&cards)
	// SELECT * FROM cards WHERE user_id = id OR user_id=nil;
	return cards
}
func (r *repository) GetCardById(userId int) models.Card {
	var card models.Card
	r.db.First(&card, "id = ?", userId)
	return card
}

func (r *repository) DeleteCard(cardId int) error {
	var cards []models.Card
	err := r.db.Where("id = ?", cardId).Find(&cards).Error
	return err
}
func (r *repository) EditCard(cardId int, cardModel models.Card) (models.Card, error) {
	var card models.Card
	err := r.db.Model(&card).Where("id = ?", cardId).Updates(cardModel).Error
	return card, err
}
