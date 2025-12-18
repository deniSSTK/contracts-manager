package context

import (
	"contracts-manager/internal/domain/auth"

	"github.com/gin-gonic/gin"
)

type ContextValue string

const (
	AuthUser ContextValue = "auth_user"
)

func SetContextValue(c *gin.Context, key ContextValue, value any) {
	c.Set(string(key), value)
}

func getContextValue[T any](c *gin.Context, key string) (T, error) {
	var zero T

	value, exists := c.Get(key)
	if !exists {
		return zero, ErrContextValueNotFound
	}

	typedValue, ok := value.(T)
	if !ok {
		return zero, ErrInvalidContextType
	}

	return typedValue, nil
}

func GetAuthUser(c *gin.Context) (*auth.AuthUser, error) {
	return getContextValue[*auth.AuthUser](c, string(AuthUser))
}
