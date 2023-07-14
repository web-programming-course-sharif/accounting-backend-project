package initializers

import (
	"accounting-project/routes"
	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	routes.AuthRoutes(e)
	routes.UserRoutes(e)
}
