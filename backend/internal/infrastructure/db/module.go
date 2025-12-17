package db

import (
	"contracts-manager/internal/infrastructure/db/repositories"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		repositories.NewBaseRepository,
		repositories.NewUserRepository,
		repositories.NewPersonRepository,
	),
)
