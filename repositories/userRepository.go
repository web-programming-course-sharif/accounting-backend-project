package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email string) models.User
	UpdateTokenByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUserByEmail(email string) models.User {
	var user models.User
	r.db.First(&user, "email = ?", email)
	return user
}
func (r *repository) UpdateTokenByEmail(email string) (models.User, error) {
	return models.User{}, nil
}
func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
