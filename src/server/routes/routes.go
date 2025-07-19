package routes

import (
	// Generate automatically the swagger docs

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"

	_ "go-api/.internal/docs"
	"go-api/src/handlers"
)

// Params defines the dependencies for the routes module.
type Params struct {
	fx.In

	Echo               *echo.Echo
	HealthcheckHandler handlers.HealthcheckHandler
}

// Register registers the routes for the API.
func Register(p Params) {
	p.Echo.GET("/", p.HealthcheckHandler.GetAPIStatus)

	p.Echo.GET("/swagger/*any", echoSwagger.WrapHandler)
}
