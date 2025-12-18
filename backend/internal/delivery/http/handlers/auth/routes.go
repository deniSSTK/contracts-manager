package authhandler

import (
	"contracts-manager/internal/delivery/http/middleware"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	handler        *Handler
	authMiddleware *middleware.AuthMiddleware
}

func NewRoutes(
	handler *Handler,
	authMiddleware *middleware.AuthMiddleware,
) *Routes {
	return &Routes{
		handler,
		authMiddleware,
	}
}

func (r *Routes) RegisterRoutes(routes *gin.RouterGroup) {
	group := routes.Group("/auth")

	group.GET("/user", r.authMiddleware.Middleware(), r.handler.GetAuthUser)
	group.GET("/refresh/access", r.authMiddleware.Middleware(), r.handler.RefreshAccessToken)

	group.POST("/login", r.handler.Login)
	group.POST("/signup", r.handler.Signup)

	group.DELETE("/logout", r.authMiddleware.Middleware(), r.handler.Logout)
}
