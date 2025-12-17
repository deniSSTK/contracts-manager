package handlers

import (
	authhandler "contracts-manager/internal/delivery/http/handlers/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		authhandler.NewHandler,
		authhandler.NewRoutes,
	),

	fx.Invoke(registerRoutes),
)

func registerRoutes(
	routes *gin.RouterGroup,
	authRoutes *authhandler.Routes,
) {
	authRoutes.RegisterRoutes(routes)
}
