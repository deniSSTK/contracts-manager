package context

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func getToken(c *gin.Context, tokenType string) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", ErrMissingAuthorizationHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)

	if len(parts) != 2 {
		return "", ErrInvalidAuthorizationHeader
	}

	if parts[0] != tokenType {
		return "", ErrUnexpectedTokenType
	}

	return parts[1], nil
}

func GetAccessToken(c *gin.Context) (string, error) {
	return getToken(c, "Bearer")
}
