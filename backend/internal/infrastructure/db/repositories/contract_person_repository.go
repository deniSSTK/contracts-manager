package repositories

import (
	"context"
	"contracts-manager/internal/domain/contract"
	"contracts-manager/internal/infrastructure/db/models"

	"github.com/google/uuid"
)

func (r *ContractRepository) AddPerson(
	ctx context.Context,
	dto contract.AddPersonDTO,
) (*models.ContractPerson, error) {
	cp := models.ContractPerson{
		ContractID: dto.ContractID,
		PersonID:   dto.PersonID,
		Role:       dto.Role,
	}

	if err := r.db.WithContext(ctx).Create(&cp).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).
		Preload("Contract").
		Preload("Person").
		First(&cp, "contract_id = ? AND person_id = ?", cp.ContractID, cp.PersonID).Error; err != nil {
		return nil, err
	}

	return &cp, nil
}

func (r *ContractRepository) RemovePerson(
	ctx context.Context,
	contractID, personID uuid.UUID,
) error {
	return r.db.WithContext(ctx).
		Where("contract_id = ? AND person_id = ?", contractID, personID).
		Unscoped().
		Delete(&models.ContractPerson{}).
		Error
}

func (r *ContractRepository) GetPersons(
	ctx context.Context,
	contractID uuid.UUID,
) ([]models.Person, error) {
	var persons []models.Person

	err := r.db.WithContext(ctx).
		Model(&models.Person{}).
		Joins("JOIN contract_people cp ON cp.person_id = people.id").
		Where("cp.contract_id = ?", contractID).
		Find(&persons).Error

	return persons, err
}

func (r *ContractRepository) GetContractsByPerson(
	ctx context.Context,
	personID uuid.UUID,
) ([]models.Contract, error) {
	var contracts []models.Contract

	err := r.db.WithContext(ctx).
		Model(&models.Contract{}).
		Joins("JOIN contract_people cp ON cp.contract_id = contracts.id").
		Where("cp.person_id = ?", personID).
		Find(&contracts).Error

	return contracts, err
}

func (r *ContractRepository) IsPersonInContract(
	ctx context.Context,
	personID, contractID uuid.UUID,
) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&models.ContractPerson{}).
		Where("person_id = ? AND contract_id = ?", personID, contractID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
