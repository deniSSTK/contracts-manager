package handlers

import (
	authhandler "contracts-manager/internal/delivery/http/handlers/auth"
	contracthandler "contracts-manager/internal/delivery/http/handlers/contract"
	personhandler "contracts-manager/internal/delivery/http/handlers/person"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		authhandler.NewHandler,
		authhandler.NewRoutes,

		personhandler.NewHandler,
		personhandler.NewRoutes,

		contracthandler.NewHandler,
		contracthandler.NewRoutes,
	),

	fx.Invoke(registerRoutes),
)

func registerRoutes(
	routes *gin.RouterGroup,
	authRoutes *authhandler.Routes,
	personRoutes *personhandler.Routes,
	contractRoutes *contracthandler.Routes,
) {
	authRoutes.RegisterRoutes(routes)
	personRoutes.RegisterRoutes(routes)
	contractRoutes.RegisterRoutes(routes)
}
