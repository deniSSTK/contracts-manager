package auth

import (
	"contracts-manager/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

type LoginDTO struct {
	UsernameOrEmail string `binding:"required,min=5" json:"usernameOrEmail"`
	Password        string `binding:"required,min=8" json:"password"`
}

type SignupDTO struct {
	Username string `binding:"required,min=5,max=50" json:"username"`
	Email    string `binding:"required,email" json:"email"`
	Password string `binding:"required,min=8" json:"password"`
}

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	Exp         int64  `json:"exp"`
}

type Filter struct {
	Username *string
	Email    *string
	Type     *string
	Page     int
	Limit    int
}

type UpdateDTO struct {
	Username *string          `json:"username"`
	Email    *string          `json:"email"`
	Type     *models.UserType `json:"type"`
	PersonID **uuid.UUID      `json:"personId"`
}
