package contracthandler

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
	group := routes.Group("/contract")

	group.GET("/:id", r.authMiddleware.AdminOnly(), r.handler.Get)
	group.GET("/:id/persons", r.authMiddleware.AdminOnly(), r.handler.GetPersons)
	group.GET("/person/:id", r.authMiddleware.AdminOnly(), r.handler.GetContractsByPerson)

	group.PUT("/:id", r.authMiddleware.AdminOnly(), r.handler.Update)

	group.POST("/", r.authMiddleware.AdminOnly(), r.handler.Create)
	group.POST("/person", r.authMiddleware.AdminOnly(), r.handler.AddPerson)

	group.DELETE("/:id", r.authMiddleware.AdminOnly(), r.handler.Delete)
	group.DELETE("/:id/person/:personId", r.authMiddleware.AdminOnly(), r.handler.RemovePerson)
}
