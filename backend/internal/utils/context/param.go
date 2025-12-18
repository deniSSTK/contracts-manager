package context

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUIDFromParam(c *gin.Context, key string) (uuid.UUID, error) {
	idStr := c.Param(key)
	id, err := uuid.Parse(idStr)

	if err != nil {
		return uuid.Nil, err
	}

	if id == uuid.Nil {
		return uuid.Nil, ErrNilIDAfterParseParam
	}

	return id, nil
}

func GetIdFromParam(c *gin.Context) (uuid.UUID, error) {
	return GetUUIDFromParam(c, "id")
}
