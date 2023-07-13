package authDto

type SignUpResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
