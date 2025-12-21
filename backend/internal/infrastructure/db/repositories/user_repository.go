package repositories

import (
	"context"
	"contracts-manager/internal/delivery/http/dto"
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/models"
	"strings"

	"github.com/google/uuid"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(baseRepository BaseRepository) *UserRepository {
	return &UserRepository{baseRepository}
}

func (r *UserRepository) Create(
	ctx context.Context,
	dto auth.SignupDTO,
	passwordHash string,
) (uuid.UUID, error) {
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

func (r *UserRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.User, error) {
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

func (r *UserRepository) GetUserIDByUsernameOrEmail(
	ctx context.Context,
	usernameOrEmail string,
) (uuid.UUID, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).
		First(&user).Error

	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (r *UserRepository) CheckEmailExists(
	ctx context.Context,
	email string,
) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepository) List(
	ctx context.Context,
	filter auth.Filter,
) (*dto.ListResult[models.User], error) {
	var users []models.User
	var total int64

	query := r.db.WithContext(ctx).Model(&models.User{})

	if filter.Username != nil && *filter.Username != "" {
		query = query.Where(
			"LOWER(username) LIKE ?",
			"%"+strings.ToLower(*filter.Username)+"%",
		)
	}

	if filter.Email != nil && *filter.Email != "" {
		query = query.Where(
			"LOWER(email) LIKE ?",
			"%"+strings.ToLower(*filter.Email)+"%",
		)
	}

	if filter.Type != nil && *filter.Type != "" {
		query = query.Where(
			"LOWER(type) LIKE ?",
			"%"+strings.ToLower(*filter.Type)+"%",
		)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	if filter.Page <= 0 {
		filter.Page = 1
	}
	if filter.Limit <= 0 {
		filter.Limit = 20
	}

	offset := (filter.Page - 1) * filter.Limit
	if err := query.Limit(filter.Limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	return &dto.ListResult[models.User]{
		Data:  users,
		Page:  filter.Page,
		Limit: filter.Limit,
		Total: total,
	}, nil
}

func (r *UserRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	data map[string]interface{},
) (*models.User, error) {
	if err := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetByID(ctx, id)
}
