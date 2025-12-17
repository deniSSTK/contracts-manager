package infrastructure

import (
	"contracts-manager/internal/infrastructure/config"
	"contracts-manager/internal/infrastructure/db"
	"contracts-manager/internal/infrastructure/logger"
	"contracts-manager/internal/infrastructure/token"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewConfig,
		logger.NewLogger,
		db.NewDB,
		token.NewJWTProvider,
	),

	db.Module,
)
