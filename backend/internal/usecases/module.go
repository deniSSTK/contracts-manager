package usecases

import (
	authusecase "contracts-manager/internal/usecases/auth"
	userusecase "contracts-manager/internal/usecases/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		authusecase.NewUsecase,
		userusecase.NewUsecase,
	),
)
