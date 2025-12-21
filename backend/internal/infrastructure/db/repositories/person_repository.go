package repositories

import (
	"context"
	"contracts-manager/internal/delivery/http/dto"
	"contracts-manager/internal/domain/person"
	"contracts-manager/internal/infrastructure/db/models"
	"strings"

	"github.com/google/uuid"
)

type PersonRepository struct {
	BaseRepository
}

func NewPersonRepository(baseRepository BaseRepository) *PersonRepository {
	return &PersonRepository{baseRepository}
}

func (r *PersonRepository) Create(
	ctx context.Context,
	dto person.CreateDTO,
) error {
	newPerson := models.Person{
		Type:  dto.Type,
		Name:  dto.Name,
		Code:  dto.Code,
		Email: dto.Email,
		Phone: dto.Phone,
	}

	return r.db.WithContext(ctx).Create(&newPerson).Error
}

func (r *PersonRepository) GetByID(
	ctx context.Context,
	personID uuid.UUID,
) (*models.Person, error) {
	var p models.Person
	if err := r.db.WithContext(ctx).First(&p, "id = ?", personID).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *PersonRepository) Update(
	ctx context.Context,
	personID uuid.UUID,
	data map[string]interface{},
) (*models.Person, error) {
	if err := r.db.WithContext(ctx).
		Model(&models.Person{}).
		Where("id = ?", personID).
		Updates(data).Error; err != nil {
		return nil, err
	}

	return r.GetByID(ctx, personID)
}

func (r *PersonRepository) Delete(ctx context.Context, personID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.Person{}, "id = ?", personID).Error
}

func (r *PersonRepository) List(
	ctx context.Context,
	filter person.Filter,
) (*dto.ListResult[models.Person], error) {
	var persons []models.Person
	var total int64

	query := r.db.WithContext(ctx).Model(&models.Person{})

	if filter.Code != nil && *filter.Code != "" {
		query = query.Where(
			"LOWER(code) LIKE ?",
			"%"+strings.ToLower(*filter.Code)+"%",
		)
	}

	if filter.Type != nil && *filter.Type != "" {
		query = query.Where(
			"LOWER(type) LIKE ?",
			"%"+strings.ToLower(*filter.Type)+"%",
		)
	}

	if filter.Name != nil && *filter.Name != "" {
		query = query.Where(
			"LOWER(name) LIKE ?",
			"%"+strings.ToLower(*filter.Name)+"%",
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
	if err := query.Limit(filter.Limit).Offset(offset).Find(&persons).Error; err != nil {
		return nil, err
	}

	return &dto.ListResult[models.Person]{
		Data:  persons,
		Page:  filter.Page,
		Limit: filter.Limit,
		Total: total,
	}, nil
}

func (r *ContractRepository) GetContractsByID(
	ctx context.Context,
	id uuid.UUID,
) ([]models.Contract, error) {

	var contracts []models.Contract

	err := r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Find(&contracts).
		Error

	if err != nil {
		return nil, err
	}

	return contracts, nil
}
