package auth

import "errors"

var (
	ErrFailedToParseTokenMapClaims = errors.New("failed to parse token map claims")
	ErrFailedToParseTokenUserId    = errors.New("failed to parse token userId")

	ErrRoleNotAllowed = errors.New("role not allowed")
)
