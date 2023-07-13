package authDto

import "accounting-project/models"

type SignUpResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}
