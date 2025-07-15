package services

import (
	"go.uber.org/fx"

	"go-api/src/services/healthcheck"
)

var Module = fx.Options(
	fx.Provide(
		healthcheck.New,
	),
)
