package repositories

import (
	"go.uber.org/fx"

	taskrepository "go-api/src/repositories/task-repository"
)

var Module = fx.Options(
	fx.Provide(
		taskrepository.New,
	),
)
