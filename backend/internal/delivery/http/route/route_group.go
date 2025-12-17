package route

import "github.com/gin-gonic/gin"

func NewRouteGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("/api")
}
