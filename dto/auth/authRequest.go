package authDto

type SignUpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=9"`
	FirstName   string `json:"first_name" validate:"required,min=3"`
	LastName    string `json:"last_name" validate:"required,min=3"`
	Password    string `json:"password" validate:"required,min=8"`
	Code        string `json:"code"`
}
type VerifyRequest struct {
	PhoneNumber string `json:"phone_number" validation:"required"`
	Code        string `json:"code" validation:"required"`
}
type ResendRequest struct {
	PhoneNumber string `json:"phone_number" validation:"required"`
}
type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=9"`
	Password    string `json:"password" validate:"required,min=8"`
}
type ForgotRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=9"`
}
