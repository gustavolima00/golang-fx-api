package config

import (
	"go.uber.org/fx"

	"go-api/src/config/config"
)

var Module = fx.Options(
	fx.Provide(
		config.New,
	),
)
