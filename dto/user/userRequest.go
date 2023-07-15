package user

type EditProfileStatusRequest struct {
	IsPublic bool `json:"is_public" validation:"required"`
}
type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password" validate:"required,min=8"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8"`
}
