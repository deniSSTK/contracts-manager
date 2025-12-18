package auth

import (
	"contracts-manager/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID   uuid.UUID       `json:"id"`
	Type models.UserType `json:"type"`
}
