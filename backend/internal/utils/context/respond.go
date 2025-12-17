package context

import (
	"github.com/gin-gonic/gin"
)

func RespondWithValue(c *gin.Context, code int, body any) {
	c.JSON(code, body)
}

func RespondVoid(c *gin.Context, code int) {
	c.Status(code)
}

func RespondError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
}
