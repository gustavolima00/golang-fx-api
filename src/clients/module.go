package clients

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"clients",
	fx.Provide(
		NewPostgresClient,
	),
)
