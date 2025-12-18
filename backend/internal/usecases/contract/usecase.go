package contractusecase

import (
	"context"
	"contracts-manager/internal/domain/contract"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/db/repositories"

	"github.com/google/uuid"
)

type Usecase struct {
	contractRepo *repositories.ContractRepository
}

func NewUsecase(contractRepo *repositories.ContractRepository) *Usecase {
	return &Usecase{contractRepo}
}

func (uc *Usecase) Create(
	ctx context.Context,
	dto contract.CreateDTO,
) (*models.Contract, error) {
	return uc.contractRepo.Create(ctx, dto)
}

func (uc *Usecase) Update(
	ctx context.Context,
	id uuid.UUID,
	dto contract.UpdateDTO,
) (*models.Contract, error) {
	data := map[string]interface{}{}

	if dto.Number != nil {
		data["number"] = *dto.Number
	}
	if dto.Title != nil {
		data["title"] = *dto.Title
	}
	if dto.Description != nil {
		data["description"] = *dto.Description
	}
	if dto.StartDate != nil {
		data["start_date"] = *dto.StartDate
	}
	if dto.EndDate != nil {
		data["end_date"] = *dto.EndDate
	}

	return uc.contractRepo.Update(ctx, id, data)
}

func (uc *Usecase) GetByID(
	ctx context.Context,
	id uuid.UUID,
) (*models.Contract, error) {
	return uc.contractRepo.GetByID(ctx, id)
}

func (uc *Usecase) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {
	return uc.contractRepo.Delete(ctx, id)
}
