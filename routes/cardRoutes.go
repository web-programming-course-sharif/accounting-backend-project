package routes

import (
	"accounting-project/handlers"
	"accounting-project/pkg/middleware"
	"accounting-project/pkg/postgres"
	"accounting-project/repositories"
	"github.com/labstack/echo/v4"
)

func CardRoutes(e *echo.Group) {
	cardRepository := repositories.RepositoryCard(postgres.DB)
	h := handlers.HandlerCard(cardRepository)

	e.POST("/createCard", middleware.Auth(h.CreateCard))
	e.POST("/deleteCard", middleware.Auth(h.DeleteCard))
	e.POST("/editCard", middleware.Auth(h.EditCard))
	e.POST("/getAllCards", middleware.Auth(h.GetAllCards))
}
