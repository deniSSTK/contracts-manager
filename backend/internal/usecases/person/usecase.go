package personusecase

import (
	"context"
	"contracts-manager/internal/domain/person"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/db/repositories"
	fileusecase "contracts-manager/internal/usecases/file"
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type Usecase struct {
	personRepo  *repositories.PersonRepository
	fileUsecase *fileusecase.Usecase[models.Person, person.CreateDTO]
}

func NewUsecase(
	personRepo *repositories.PersonRepository,
) *Usecase {
	return &Usecase{
		personRepo: personRepo,
		fileUsecase: fileusecase.NewUsecase[models.Person, person.CreateDTO](
			repositories.NewStreamRepository[models.Person](personRepo.BaseRepository),
		),
	}
}

func (uc *Usecase) Create(ctx context.Context, dto person.CreateDTO) (uuid.UUID, error) {
	return uc.personRepo.Create(ctx, dto)
}

func (uc *Usecase) GetByID(ctx context.Context, personID uuid.UUID) (*models.Person, error) {
	return uc.personRepo.GetByID(ctx, personID)
}

func (uc *Usecase) Update(
	ctx context.Context,
	personID uuid.UUID,
	dto person.UpdateDTO,
) (*models.Person, error) {
	data := map[string]interface{}{}

	if dto.Type != nil {
		data["type"] = *dto.Type
	}
	if dto.Name != nil {
		data["name"] = *dto.Name
	}
	if dto.Code != nil {
		data["code"] = *dto.Code
	}
	if dto.Email != nil {
		data["email"] = *dto.Email
	}
	if dto.Phone != nil {
		data["phone"] = *dto.Phone
	}

	return uc.personRepo.Update(ctx, personID, data)
}

func (uc *Usecase) Delete(ctx context.Context, personID uuid.UUID) error {
	return uc.personRepo.Delete(ctx, personID)
}

func (uc *Usecase) List(ctx context.Context, filter person.PersonFilter) (*person.PersonListResult, error) {
	return uc.personRepo.List(ctx, filter)
}

func (uc *Usecase) ImportJSON(ctx context.Context, reader io.Reader) (int, []string) {
	return uc.fileUsecase.ImportJSON(
		ctx,
		reader,
		func(ctx context.Context, dto person.CreateDTO) error {
			_, err := uc.personRepo.Create(ctx, dto)
			return err
		},
		func(decoder *json.Decoder) (person.CreateDTO, error) {
			var dto person.CreateDTO
			err := decoder.Decode(&dto)
			return dto, err
		},
	)
}

func (uc *Usecase) ImportCSV(ctx context.Context, reader io.Reader) (int, []string) {
	return uc.fileUsecase.ImportCSV(
		ctx,
		reader,
		func(ctx context.Context, dto person.CreateDTO) error {
			_, err := uc.personRepo.Create(ctx, dto)
			return err
		},
		func(fileHeaders, record []string) (person.CreateDTO, error) {
			dto := person.CreateDTO{}
			for i, h := range fileHeaders {
				switch h {
				case "type":
					dto.Type = models.PersonType(record[i])
				case "name":
					dto.Name = record[i]
				case "code":
					dto.Code = record[i]
				case "email":
					if record[i] != "" {
						dto.Email = &record[i]
					}
				case "phone":
					if record[i] != "" {
						dto.Phone = &record[i]
					}
				}
			}
			return dto, nil
		},
	)
}

func (uc *Usecase) ExportCSV(ctx context.Context, w io.Writer) error {
	headers := []string{"id", "type", "name", "code", "email", "phone"}

	return uc.fileUsecase.ExportCSV(ctx, w, headers, func(p models.Person) []string {
		record := []string{
			p.ID.String(),
			string(p.Type),
			p.Name,
			p.Code,
		}

		if p.Email != nil {
			record = append(record, *p.Email)
		} else {
			record = append(record, "")
		}

		if p.Phone != nil {
			record = append(record, *p.Phone)
		} else {
			record = append(record, "")
		}

		return record
	})
}

func (uc *Usecase) ExportJSON(ctx context.Context, w io.Writer) error {
	return uc.fileUsecase.ExportJSON(ctx, w)
}
