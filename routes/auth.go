package routes

import (
	"accounting-project/handlers"
	"accounting-project/pkg/middleware"
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
	e.POST("/forgot", h.Forgot)
	e.POST("/resendCode", h.Resend)
	e.GET("/myAccount", middleware.Auth(h.CheckAuth))
}
