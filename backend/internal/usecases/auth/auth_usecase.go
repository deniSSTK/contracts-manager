package authusecase

import (
	"context"
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/repositories"
	"contracts-manager/internal/utils"

	"github.com/google/uuid"
)

type Usecase struct {
	userRepo *repositories.UserRepository
}

func NewUsecase(userRepo *repositories.UserRepository) *Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}

func (uc *Usecase) GetAuthUser(ctx context.Context, id uuid.UUID) (*auth.AuthUser, error) {
	user, err := uc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return auth.ParseAuthUserFromUserModel(user), nil
}

func (uc *Usecase) Login(ctx context.Context, dto auth.LoginDTO) (uuid.UUID, error) {
	passwordHash, err := uc.userRepo.GetPasswordHashByUsernameOrEmail(ctx, dto.UsernameOrEmail)
	if err != nil {
		return uuid.Nil, err
	}

	if valid := utils.CheckPassword([]byte(passwordHash), dto.Password); !valid {
		return uuid.Nil, ErrIncorrectPassword
	}

	return uc.userRepo.GetUserIDByUsernameOrEmail(ctx, dto.UsernameOrEmail)
}

func (uc *Usecase) Signup(ctx context.Context, dto auth.SignupDTO) (uuid.UUID, error) {
	emailExists, err := uc.userRepo.CheckEmailExists(ctx, dto.Email)
	if err != nil {
		return uuid.Nil, err
	}

	if emailExists {
		return uuid.Nil, ErrEmailAlreadyExists
	}

	passwordHash, err := utils.HashPassword(dto.Password)
	if err != nil {
		return uuid.Nil, err
	}

	return uc.userRepo.Insert(ctx, dto, passwordHash)
}
