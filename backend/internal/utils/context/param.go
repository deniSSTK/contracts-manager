package context

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUUIDFromParam(c *gin.Context) (uuid.UUID, error) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)

	if err != nil {
		return uuid.Nil, err
	}

	if id == uuid.Nil {
		return uuid.Nil, ErrNilIDAfterParseParam
	}

	return id, nil
}
