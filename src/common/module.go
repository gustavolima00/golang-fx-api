package common

import (
	"go-api/src/common/config"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"common",
	fx.Provide(
		config.NewConfig,
	),
)
