package routes

import (
	"accounting-project/handlers"
	"accounting-project/pkg/middleware"
	"accounting-project/pkg/postgres"
	"accounting-project/repositories"
	"github.com/labstack/echo/v4"
)

func BankRoutes(e *echo.Group) {
	bankRepository := repositories.RepositoryBank(postgres.DB)
	h := handlers.HandlerBank(bankRepository)

	e.POST("/getAllBank", middleware.Auth(h.GetAllBanks))
}
