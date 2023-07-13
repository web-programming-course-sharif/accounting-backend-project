package routes

import (
	"accounting-project/handlers"
	"accounting-project/pkg/postgres"
	"accounting-project/repositories"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(postgres.DB)
	h := handlers.HandlerUser(userRepository)

	e.POST("/signUp", h.SignUp)
	e.POST("/login", h.Login)
	e.POST("/verify", h.Verify)
	e.POST("/myAccount", h.Account)
	e.POST("/forgot", h.Forgot)
	//e.GET("/check-auth", middleware.Auth(h.CheckAuth)) // add this code
}
