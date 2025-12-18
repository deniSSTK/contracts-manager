package usecases

import (
	authusecase "contracts-manager/internal/usecases/auth"
	contractusecase "contracts-manager/internal/usecases/contract"
	personusecase "contracts-manager/internal/usecases/person"
	userusecase "contracts-manager/internal/usecases/user"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		authusecase.NewUsecase,
		userusecase.NewUsecase,
		personusecase.NewUsecase,
		contractusecase.NewUsecase,
	),
)
