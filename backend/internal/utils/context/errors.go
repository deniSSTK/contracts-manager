package context

import "errors"

var (
	ErrNilIdAfterParseParam = errors.New("nil id after parse param")

	ErrContextValueNotFound = errors.New("value not found in contex")
	ErrInvalidContextType   = errors.New("invalid type for context value")

	ErrMissingAuthorizationHeader = errors.New("authorization header is missing")
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	ErrUnexpectedTokenType        = errors.New("unexpected token type")
)
