package authusecase

import (
	"context"
	"contracts-manager/internal/delivery/http/dto"
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/db/repositories"
	contractusecase "contracts-manager/internal/usecases/contract"
	"contracts-manager/internal/utils"

	"github.com/google/uuid"
)

type Usecase struct {
	userRepo   *repositories.UserRepository
	contractUC *contractusecase.Usecase
}

func NewUsecase(
	userRepo *repositories.UserRepository,
	contractUC *contractusecase.Usecase,
) *Usecase {
	return &Usecase{
		userRepo,
		contractUC,
	}
}

func (uc *Usecase) GetAuthUser(
	ctx context.Context,
	id uuid.UUID,
) (*auth.AuthUser, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return auth.ParseAuthUserFromUserModel(user), nil
}

func (uc *Usecase) Login(
	ctx context.Context,
	dto auth.LoginDTO,
) (uuid.UUID, error) {
	passwordHash, err := uc.userRepo.GetPasswordHashByUsernameOrEmail(ctx, dto.UsernameOrEmail)
	if err != nil {
		return uuid.Nil, err
	}

	if valid := utils.CheckPassword([]byte(passwordHash), dto.Password); !valid {
		return uuid.Nil, ErrIncorrectPassword
	}

	return uc.userRepo.GetUserIDByUsernameOrEmail(ctx, dto.UsernameOrEmail)
}

func (uc *Usecase) Signup(
	ctx context.Context,
	dto auth.SignupDTO,
) (uuid.UUID, error) {
	if dto.Type == nil {
		usertype := models.UserTypeRegular
		dto.Type = &usertype
	}

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

	return uc.userRepo.Create(ctx, dto, passwordHash)
}

func (uc *Usecase) GetByID(ctx context.Context, userId uuid.UUID) (*models.User, error) {
	return uc.userRepo.GetByID(ctx, userId)
}

func (uc *Usecase) List(ctx context.Context, filter auth.Filter) (*dto.ListResult[models.User], error) {
	return uc.userRepo.List(ctx, filter)
}

func (uc *Usecase) Update(
	ctx context.Context,
	id uuid.UUID,
	dto auth.UpdateDTO,
) (*models.User, error) {
	data := map[string]interface{}{}

	if dto.Username != nil {
		data["username"] = *dto.Username
	}
	if dto.Email != nil {
		data["email"] = *dto.Email
	}
	if dto.Type != nil {
		data["type"] = *dto.Type
	}
	if dto.PersonID != nil {
		data["personId"] = *dto.PersonID
	}

	return uc.userRepo.Update(ctx, id, data)
}

func (uc *Usecase) GetContractsByID(
	ctx context.Context,
	id uuid.UUID,
) ([]models.Contract, error) {
	user, err := uc.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.PersonID == nil {
		return nil, ErrNilPersonID
	}

	return uc.contractUC.GetContractsByPerson(ctx, *user.PersonID)
}
