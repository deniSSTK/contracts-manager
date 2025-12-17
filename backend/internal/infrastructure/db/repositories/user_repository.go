package repositories

import (
	"context"
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(baseRepository BaseRepository) *UserRepository {
	return &UserRepository{
		BaseRepository: baseRepository,
	}
}

func (r *UserRepository) Insert(ctx context.Context, dto auth.SignupDTO, passwordHash string) (uuid.UUID, error) {
	newUser := models.User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: passwordHash,
		Type:         models.UserTypeRegular,
	}

	if err := r.db.WithContext(ctx).Create(&newUser).Error; err != nil {
		return uuid.Nil, err
	}

	return newUser.ID, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetPasswordHashByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (string, error) {
	var passwordHash string

	err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Select("password_hash").
		Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).
		First(&passwordHash).Error

	if err != nil {
		return "", err
	}

	return passwordHash, nil
}

func (r *UserRepository) GetUserIDByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (uuid.UUID, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).
		First(&user).Error

	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (r *UserRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
