package server

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	"go-api/src/common/config"
	"go-api/src/services"
)

// Params defines the dependencies that the server module needs
type Params struct {
	fx.In

	Lifecycle          fx.Lifecycle
	HealthcheckService services.HealthcheckService
	Config             *config.Config
}

// NewServer returns a pointer to Server
func NewServer(p Params) *echo.Echo {
	e := echo.New()

	p.Lifecycle.Append(fx.Hook{
		OnStart: onStart(e, p),
		OnStop:  onStop(e, p),
	})

	return e
}

func onStart(e *echo.Echo, p Params) func(context.Context) error {
	return func(ctx context.Context) error {
		p.HealthcheckService.SetOnlineSince(time.Now())
		go e.Start(":" + p.Config.Port)
		return nil
	}
}

func onStop(e *echo.Echo, p Params) func(context.Context) error {
	return func(ctx context.Context) error {
		log.Println("Stopping server")
		return e.Shutdown(ctx)
	}
}
