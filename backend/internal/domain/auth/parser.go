package auth

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/google/uuid"
)

func ParseUserIDFromJWTClaims(claims jwt.Claims) (uuid.UUID, error) {
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, ErrFailedToParseTokenMapClaims
	}

	userIDValue, ok := mapClaims["userID"]
	if !ok {
		return uuid.Nil, ErrFailedToParseTokenUserID
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		return uuid.Nil, ErrFailedToParseTokenUserID
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
