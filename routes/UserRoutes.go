package routes

import (
	"accounting-project/handlers"
	"accounting-project/pkg/middleware"
	"accounting-project/pkg/postgres"
	"accounting-project/repositories"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(postgres.DB)
	h := handlers.HandlerUser(userRepository)

	e.POST("/editProfileStatus", middleware.Auth(h.EditProfileStatus))
	e.POST("/changePassword", middleware.Auth(h.ChangePassword))
	e.POST("/setPhoto", middleware.UploadFile(h.ChangePassword))
	e.POST("/changeProfile", middleware.UploadFile(h.EditProfile))
}
