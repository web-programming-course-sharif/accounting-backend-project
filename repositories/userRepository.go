package repositories

import (
	"accounting-project/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByPhoneNumber(phoneNumber string) models.User
	UpdateTokenByPhoneNumber(phoneNumber string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	CheckAuth(id int) models.User
	EditProfileStatus(id int, status bool) models.User
	ChangePassword(id int, password string) models.User
	ChangeProfile(id int, user models.User) models.User
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUserByPhoneNumber(phoneNumber string) models.User {
	var user models.User
	r.db.First(&user, "phone_number = ?", phoneNumber)
	return user
}
func (r *repository) UpdateTokenByPhoneNumber(phoneNumber string) (models.User, error) {
	return models.User{}, nil
}
func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
func (r *repository) CheckAuth(id int) models.User {
	var user models.User
	r.db.First(&user, "id = ?", id)
	return user
}
func (r *repository) EditProfileStatus(id int, status bool) models.User {
	var user models.User
	r.db.Model(&user).Where("id = ?", id).Update("is_public", status)
	r.db.First(&user, "id = ?", id)
	return user
}
func (r *repository) ChangePassword(id int, password string) models.User {
	var user models.User
	r.db.Model(&user).Where("id = ?", id).Update("password", password)
	r.db.First(&user, "id = ?", id)
	return user
}
func (r *repository) ChangeProfile(id int, userModel models.User) models.User {
	var user models.User
	r.db.Model(&user).Where("id = ?", id).Updates(userModel)
	return user
}
