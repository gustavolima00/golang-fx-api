package server

import (
	_ "go-api/.internal/docs" // Generate automatically the swagger docs
	"go-api/src/handlers/auth"
	"go-api/src/handlers/healthcheck"
	"go-api/src/server/middlewares"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/fx"
)

// Params defines the dependencies for the routes module.
type RegisterRoutesParams struct {
	fx.In

	Echo        *echo.Echo
	Healthcheck healthcheck.Handler
	AuthHandler auth.AuthHandler
	Middlewares middlewares.Middlewares
}

// RegisterRoutes registers the routes for the API.
func RegisterRoutes(p RegisterRoutesParams) {
	// Base routes
	p.Echo.GET("/", p.Healthcheck.GetAPIStatus)
	p.Echo.GET("/swagger/*any", echoSwagger.WrapHandler)

	// Authentication routes
	authGroup := p.Echo.Group("/auth")
	{
		authGroup.POST("/login", p.AuthHandler.CreateSession)
		authGroup.POST("/refresh", p.AuthHandler.UpdateSession)
		authGroup.POST("/logout", p.AuthHandler.FinishSession)
		authGroup.GET("/user", p.AuthHandler.GetUser, p.Middlewares.AuthMiddleware())
	}
}
