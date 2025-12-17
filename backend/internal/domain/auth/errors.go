package auth

import "errors"

var (
	ErrFailedToParseTokenMapClaims = errors.New("failed to parse token map claims")
	ErrFailedToParseTokenUserID    = errors.New("failed to parse token userID")

	ErrRoleNotAllowed = errors.New("role not allowed")
)
