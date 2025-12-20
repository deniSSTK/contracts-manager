package contractusecase

import (
	"context"
	"contracts-manager/internal/delivery/http/dto"
	"contracts-manager/internal/domain/contract"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/db/repositories"
	fileusecase "contracts-manager/internal/usecases/file"
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
)

type Usecase struct {
	contractRepo *repositories.ContractRepository
	fileUsecase  *fileusecase.Usecase[models.Contract, contract.CreateDTO]
}

func NewUsecase(contractRepo *repositories.ContractRepository) *Usecase {
	return &Usecase{
		contractRepo: contractRepo,
		fileUsecase: fileusecase.NewUsecase[models.Contract, contract.CreateDTO](
			repositories.NewStreamRepository[models.Contract](contractRepo.BaseRepository),
		),
	}
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

	if dto.Code != nil {
		data["code"] = *dto.Code
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

func (uc *Usecase) GetByID(ctx context.Context, id uuid.UUID) (*models.Contract, error) {
	return uc.contractRepo.GetByID(ctx, id)
}

func (uc *Usecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.contractRepo.Delete(ctx, id)
}

func (uc *Usecase) AddPerson(ctx context.Context, dto contract.AddPersonDTO) (*models.ContractPerson, error) {
	return uc.contractRepo.AddPerson(ctx, dto)
}

func (uc *Usecase) RemovePerson(ctx context.Context, contractID, personID uuid.UUID) error {
	return uc.contractRepo.RemovePerson(ctx, contractID, personID)
}

func (uc *Usecase) GetPersons(ctx context.Context, contractID uuid.UUID) ([]models.Person, error) {
	return uc.contractRepo.GetPersons(ctx, contractID)
}

func (uc *Usecase) GetContractsByPerson(ctx context.Context, personID uuid.UUID) ([]models.Contract, error) {
	return uc.contractRepo.GetContractsByPerson(ctx, personID)
}

func (uc *Usecase) ImportJSON(ctx context.Context, reader io.Reader) (int, []string) {
	return uc.fileUsecase.ImportJSON(
		ctx,
		reader,
		func(ctx context.Context, dto contract.CreateDTO) error {
			_, err := uc.contractRepo.Create(ctx, dto)
			return err
		},
		func(decoder *json.Decoder) (contract.CreateDTO, error) {
			var dto contract.CreateDTO
			err := decoder.Decode(&dto)
			return dto, err
		},
	)
}

func (uc *Usecase) ImportCSV(ctx context.Context, reader io.Reader) (int, []string) {
	return uc.fileUsecase.ImportCSV(
		ctx,
		reader,
		func(ctx context.Context, dto contract.CreateDTO) error {
			_, err := uc.contractRepo.Create(ctx, dto)
			return err
		},
		func(fileHeaders, record []string) (contract.CreateDTO, error) {
			dto := contract.CreateDTO{}
			for i, h := range fileHeaders {
				switch h {
				case "code":
					dto.Code = record[i]
				case "title":
					dto.Title = record[i]
				case "description":
					if record[i] != "" {
						dto.Description = &record[i]
					}
				case "startDate":
					if record[i] != "" {
						t, err := time.Parse(time.RFC3339, record[i])
						if err != nil {
							return dto, err
						}
						dto.StartDate = &t
					}
				case "endDate":
					if record[i] != "" {
						t, err := time.Parse(time.RFC3339, record[i])
						if err != nil {
							return dto, err
						}
						dto.EndDate = &t
					}
				}
			}
			return dto, nil
		},
	)
}

func (uc *Usecase) ExportCSV(ctx context.Context, w io.Writer) error {
	headers := []string{"id", "code", "title", "description", "startDate", "endDate"}

	return uc.fileUsecase.ExportCSV(ctx, w, headers, func(c models.Contract) []string {
		record := []string{
			c.ID.String(),
			c.Code,
			c.Title,
		}

		if c.Description != nil {
			record = append(record, *c.Description)
		} else {
			record = append(record, "")
		}

		if c.StartDate != nil {
			record = append(record, c.StartDate.Format(time.RFC3339))
		} else {
			record = append(record, "")
		}

		if c.EndDate != nil {
			record = append(record, c.EndDate.Format(time.RFC3339))
		} else {
			record = append(record, "")
		}

		return record
	})
}

func (uc *Usecase) ExportJSON(ctx context.Context, w io.Writer) error {
	return uc.fileUsecase.ExportJSON(ctx, w)
}

func (uc *Usecase) List(ctx context.Context, filter contract.Filter) (*dto.ListResult[models.Contract], error) {
	return uc.contractRepo.List(ctx, filter)
}
