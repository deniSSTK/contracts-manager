package personhandler

import (
	"contracts-manager/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler        *Handler
	authMiddleware *middleware.AuthMiddleware
}

func NewRoutes(handler *Handler) *Routes {
	return &Routes{
		handler: handler,
	}
}

func (r *Routes) RegisterRoutes(routes *gin.RouterGroup) {
	group := routes.Group("/person")

	group.GET("/", r.authMiddleware.AdminOnly(), r.handler.List)
	group.GET("/:id", r.authMiddleware.AdminOnly(), r.handler.GetByID)
	group.GET("/export", r.authMiddleware.AdminOnly(), r.handler.Export)

	group.PUT("/:id", r.authMiddleware.AdminOnly(), r.handler.Update)

	group.POST("/", r.authMiddleware.AdminOnly(), r.handler.Insert)
	group.POST("/import", r.authMiddleware.AdminOnly(), r.handler.Import)

	group.DELETE("/:id", r.authMiddleware.AdminOnly(), r.handler.Delete)
}
