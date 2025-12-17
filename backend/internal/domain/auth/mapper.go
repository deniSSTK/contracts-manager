package auth

import "contracts-manager/internal/infrastructure/db/models"

func ParseAuthUserFromUserModel(user *models.User) *AuthUser {
	return &AuthUser{
		ID:   user.ID,
		Type: user.Type,
	}
}
