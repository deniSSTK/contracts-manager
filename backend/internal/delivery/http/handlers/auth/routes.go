package authhandler

import "github.com/gin-gonic/gin"

type Routes struct {
	Handler *Handler
}

func NewRoutes(handler *Handler) *Routes {
	return &Routes{
		Handler: handler,
	}
}

func (r *Routes) RegisterRoutes(routes *gin.RouterGroup) {
	group := routes.Group("/auth")

	group.POST("/login", r.Handler.Login)
	group.POST("/signup", r.Handler.Signup)
}
