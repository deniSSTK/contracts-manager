package auth

import (
	"contracts-manager/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

type AuthUser struct {
	Id   uuid.UUID
	Type models.UserType
}
