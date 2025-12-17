package userusecase

import (
	"contracts-manager/internal/infrastructure/db/repositories"
)

type Usecase struct {
	userRepo *repositories.UserRepository
}

func NewUsecase(userRepo *repositories.UserRepository) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}
