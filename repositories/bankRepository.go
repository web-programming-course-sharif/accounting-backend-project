package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type BankRepository interface {
	GetAllBankWithUserId(id int) []models.Bank
}

func RepositoryBank(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GetAllBankWithUserId(id int) []models.Bank {
	var banks []models.Bank
	r.db.Where("user_id = ?", id).Or("user_id", nil).Preload("User").Find(&banks)
	// SELECT * FROM users WHERE user_id = id OR user_id=nil;
	return banks
}
