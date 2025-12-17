package authusecase

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrIncorrectPassword  = errors.New("incorrect password")
)
