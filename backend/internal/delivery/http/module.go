package http

import (
	"context"
	"contracts-manager/internal/delivery/http/handlers"
	"contracts-manager/internal/delivery/http/middleware"
	"contracts-manager/internal/delivery/http/route"
	"contracts-manager/internal/infrastructure/config"
	"contracts-manager/internal/infrastructure/logger"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Options(
	fx.Provide(
		NewEngine,
		route.NewRouteGroup,

		middleware.NewAuthMiddleware,
	),

	handlers.Module,

	fx.Invoke(runEngine),
)

func runEngine(
	lc fx.Lifecycle,
	cfg *config.Config,
	log *logger.Logger,
	engine *gin.Engine,
) {
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting HTTP server", zap.String("port", cfg.Port))

			go func() {
				if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatal(ErrFailedToRunEngine.Error(), zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping HTTP server...")
			return srv.Shutdown(ctx)
		},
	})
}
