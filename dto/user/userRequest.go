package user

type EditProfileStatusRequest struct {
	IsPublic bool `json:"is_public" validation:"required"`
}
type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password" validate:"required,min=8"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8"`
}
type EditProfileRequest struct {
	FirstName string `json:"first_name" validate:"required,min=3"`
	LastName  string `json:"last_name" validate:"required,min=3"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	State     string `json:"state"`
	City      string `json:"city"`
	ZipCode   string `json:"zip_code"`
	Address   string `json:"address"`
	About     string `json:"about"`
}
type EditSocialLinksRequest struct {
	FacebookLink  string `json:"facebook_link"`
	InstagramLink string `json:"instagram_link"`
	LinkedinLink  string `json:"linkedin_link"`
	TwitterLink   string `json:"twitter_link"`
}
