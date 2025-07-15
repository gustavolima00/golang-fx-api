package clients

import (
	"go-api/src/clients/postgres"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		postgres.NewClient,
	),
)
