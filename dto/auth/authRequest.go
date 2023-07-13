package authDto

type SignUpRequest struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=3"`
	LastName  string `json:"last_name" validate:"required,min=3"`
	Password  string `json:"password" validate:"required,min=8"`
}
type VerifyRequest struct {
	MobileNumber string `json:"mobileNumber" validation:"required"`
	Password     string `json:"password" validation:"required"`
	Code         string `json:"code" validation:"required"`
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ForgotRequest struct {
	Email string `json:"email"`
}
