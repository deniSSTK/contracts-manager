package personusecase

import (
	"context"
	"contracts-manager/internal/domain/person"
	"contracts-manager/internal/infrastructure/db/models"
	"contracts-manager/internal/infrastructure/db/repositories"
	"encoding/csv"
	"encoding/json"
	"io"
	"strconv"

	"github.com/google/uuid"
)

type Usecase struct {
	personRepo *repositories.PersonRepository
}

func NewUsecase(personRepo *repositories.PersonRepository) *Usecase {
	return &Usecase{personRepo}
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
	dto person.CreateDTO,
) error {
	updateData := make(map[string]interface{})

	if dto.Type != "" {
		updateData["type"] = dto.Type
	}
	if dto.Name != "" {
		updateData["name"] = dto.Name
	}
	if dto.Code != "" {
		updateData["code"] = dto.Code
	}
	if dto.Email != nil {
		updateData["email"] = dto.Email
	}
	if dto.Phone != nil {
		updateData["phone"] = dto.Phone
	}

	if len(updateData) == 0 {
		return nil
	}

	return uc.personRepo.Update(ctx, personID, updateData)
}

func (uc *Usecase) Delete(ctx context.Context, personID uuid.UUID) error {
	return uc.personRepo.Delete(ctx, personID)
}

func (uc *Usecase) List(ctx context.Context, filter person.PersonFilter) (*person.PersonListResult, error) {
	return uc.personRepo.List(ctx, filter)
}

func (uc *Usecase) ImportJSON(
	ctx context.Context,
	reader io.Reader,
) (int, []string) {
	decoder := json.NewDecoder(reader)
	var imported int
	var errors []string

	_, err := decoder.Token()
	if err != nil {
		return 0, []string{"invalid JSON"}
	}

	i := 0
	for decoder.More() {
		var dto person.CreateDTO
		if err = decoder.Decode(&dto); err != nil {
			errors = append(errors, "row "+strconv.Itoa(i)+": "+err.Error())
			i++
			continue
		}

		_, err = uc.personRepo.Create(ctx, dto)
		if err != nil {
			errors = append(errors, "row "+strconv.Itoa(i)+": "+err.Error())
		} else {
			imported++
		}
		i++
	}

	return imported, errors
}

func (uc *Usecase) ImportCSV(
	ctx context.Context,
	reader io.Reader,
) (int, []string) {
	r := csv.NewReader(reader)
	imported := 0
	var errors []string

	headers, err := r.Read()
	if err != nil {
		return 0, []string{"cannot read CSV header"}
	}

	rowIndex := 1
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, "row "+strconv.Itoa(rowIndex)+": "+err.Error())
			rowIndex++
			continue
		}

		dto := person.CreateDTO{}
		for i, h := range headers {
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

		_, err = uc.personRepo.Create(ctx, dto)
		if err != nil {
			errors = append(errors, "row "+strconv.Itoa(rowIndex)+": "+err.Error())
		} else {
			imported++
		}
		rowIndex++
	}

	return imported, errors
}

func (uc *Usecase) ExportCSV(
	ctx context.Context,
	w io.Writer,
) error {
	cw := csv.NewWriter(w)

	if err := cw.Write([]string{
		"id", "type", "name", "code", "email", "phone",
	}); err != nil {
		return err
	}

	err := uc.personRepo.StreamAll(ctx, func(p models.Person) error {
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

		return cw.Write(record)
	})

	cw.Flush()
	return err
}

func (uc *Usecase) ExportJSON(
	ctx context.Context,
	w io.Writer,
) error {
	enc := json.NewEncoder(w)

	if _, err := w.Write([]byte("[")); err != nil {
		return err
	}

	first := true

	err := uc.personRepo.StreamAll(ctx, func(p models.Person) error {
		if !first {
			if _, err := w.Write([]byte(",")); err != nil {
				return err
			}
		}
		first = false

		return enc.Encode(p)
	})

	if err != nil {
		return err
	}

	_, err = w.Write([]byte("]"))
	return err
}
