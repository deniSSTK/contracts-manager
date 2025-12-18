package cookie

import (
	"github.com/gin-gonic/gin"
)

type CookieNames string

const (
	RefreshToken CookieNames = "refresh_token"
)

func SetCookie(c *gin.Context, name CookieNames, value string, age int) {
	c.SetCookie(
		string(name),
		value,
		age,
		"/",
		"",
		false,
		true,
	)
}

func ClearCookie(c *gin.Context, name CookieNames) {
	c.SetCookie(
		string(name),
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}

func GetCookie(c *gin.Context, name CookieNames) (string, error) {
	value, err := c.Cookie(string(name))
	if err != nil {
		return "", err
	}

	if value == "" {
		return "", ErrMissingCookie
	}

	return value, nil
}
