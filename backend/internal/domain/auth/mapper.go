package auth

import "contracts-manager/internal/infrastructure/db/models"

func ParseAuthUserFromUserModel(user *models.User) *AuthUser {
	return &AuthUser{
		Id:   user.ID,
		Type: user.Type,
	}
}
