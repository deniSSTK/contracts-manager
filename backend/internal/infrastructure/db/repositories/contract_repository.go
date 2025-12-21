package repositories

import (
	"context"
	"contracts-manager/internal/delivery/http/dto"
	"contracts-manager/internal/domain/contract"
	"contracts-manager/internal/infrastructure/db/models"
	"strings"

	"github.com/google/uuid"
)

type ContractRepository struct {
	BaseRepository
}

func NewContractRepository(baseRepository BaseRepository) *ContractRepository {
	return &ContractRepository{baseRepository}
}

func (r *ContractRepository) Create(
	ctx context.Context,
	dto contract.CreateDTO,
) (*models.Contract, error) {
	c := models.Contract{
		Code:        dto.Code,
		Title:       dto.Title,
		Description: dto.Description,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
	}

	if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *ContractRepository) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Contract, error) {
	var c models.Contract
	if err := r.db.WithContext(ctx).First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *ContractRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	data map[string]interface{},
) (*models.Contract, error) {
	if err := r.db.WithContext(ctx).
		Model(&models.Contract{}).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetByID(ctx, id)
}

func (r *ContractRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return r.db.WithContext(ctx).
		Delete(&models.Contract{}, "id = ?", id).
		Error
}

func (r *ContractRepository) List(
	ctx context.Context,
	filter contract.Filter,
) (*dto.ListResult[models.Contract], error) {
	var contracts []models.Contract
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Contract{})

	if filter.Code != nil && *filter.Code != "" {
		query = query.Where(
			"LOWER(code) LIKE ?",
			"%"+strings.ToLower(*filter.Code)+"%",
		)
	}

	if filter.Title != nil && *filter.Title != "" {
		query = query.Where(
			"LOWER(title) LIKE ?",
			"%"+strings.ToLower(*filter.Title)+"%",
		)
	}

	if filter.Description != nil && *filter.Description != "" {
		query = query.Where(
			"LOWER(description) LIKE ?",
			"%"+strings.ToLower(*filter.Description)+"%",
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
	if err := query.Limit(filter.Limit).Offset(offset).Find(&contracts).Error; err != nil {
		return nil, err
	}

	return &dto.ListResult[models.Contract]{
		Data:  contracts,
		Page:  filter.Page,
		Limit: filter.Limit,
		Total: total,
	}, nil
}
