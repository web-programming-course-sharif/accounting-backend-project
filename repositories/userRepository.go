package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByMobileNumber(mobileNumber string) models.User
	UpdateTokenByMobileNumber(mobileNumber string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUserByMobileNumber(mobileNumber string) models.User {
	var user models.User
	r.db.First(&user, "mobile_number = ?", mobileNumber)
	return user
}
func (r *repository) UpdateTokenByMobileNumber(mobileNumber string) (models.User, error) {
	return models.User{}, nil
}
func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
