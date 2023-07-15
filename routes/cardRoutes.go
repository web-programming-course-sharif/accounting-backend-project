package routes

import (
	"accounting-project/handlers"
	"accounting-project/pkg/middleware"
	"accounting-project/pkg/postgres"
	"accounting-project/repositories"
	"github.com/labstack/echo/v4"
)

func CardRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryCard(postgres.DB)
	h := handlers.HandlerUser(userRepository)

	e.POST("/editProfileStatus", middleware.Auth(h.EditProfileStatus))
	e.POST("/changePassword", middleware.Auth(h.ChangePassword))
	e.POST("/changeProfile", middleware.Auth(middleware.UploadFile(h.EditProfile)))
	e.POST("/changeSocialLinks", middleware.Auth(h.EditSocialLinks))
}
