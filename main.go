package main

import (
	"context"
	"go-api/src/config"
	"go-api/src/handlers"
	"go-api/src/server"
	"go-api/src/services"
	"log"
	"time"

	"go.uber.org/fx"
)

// @title			Go Sample API
// @version		1.0
// @description	This is a sample API for Go using Swagger
// @termsOfService	http://swagger.io/terms/
// @contact.name
// @contact.url
// @contact.email
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @BasePath	/
func main() {
	app := fx.New(
		config.Module,
		server.Module,
		services.Module,
		handlers.Module,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
