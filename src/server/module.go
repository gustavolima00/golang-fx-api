package server

import (
	"go.uber.org/fx"

	"go-api/src/server/routes"
	"go-api/src/server/server"
)

var Module = fx.Options(
	fx.Provide(
		server.NewServer,
	),
	fx.Invoke(
		routes.Register,
	),
)
