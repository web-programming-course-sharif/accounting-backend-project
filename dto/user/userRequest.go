package user

type EditProfileStatusRequest struct {
	IsPublic bool `json:"is_public" validation:"required"`
}
