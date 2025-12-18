package fileusecase

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"strconv"

	"contracts-manager/internal/infrastructure/db/repositories"
)

type Usecase[T, TCreate any] struct {
	streamRepository *repositories.StreamRepository[T]
}

func NewUsecase[T, TCreate any](streamRepository *repositories.StreamRepository[T]) *Usecase[T, TCreate] {
	return &Usecase[T, TCreate]{streamRepository}
}

func (uc *Usecase[T, TCreate]) ImportJSON(
	ctx context.Context,
	reader io.Reader,
	createFunc func(context.Context, TCreate) error,
	decodeFunc func(*json.Decoder) (TCreate, error),
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
		dto, err := decodeFunc(decoder)
		if err != nil {
			errors = append(errors, "row "+strconv.Itoa(i)+": "+err.Error())
			i++
			continue
		}

		if err = createFunc(ctx, dto); err != nil {
			errors = append(errors, "row "+strconv.Itoa(i)+": "+err.Error())
		} else {
			imported++
		}
		i++
	}

	return imported, errors
}

func (uc *Usecase[T, TCreate]) ImportCSV(
	ctx context.Context,
	reader io.Reader,
	createFunc func(context.Context, TCreate) error,
	convertFunc func([]string, []string) (TCreate, error),
) (int, []string) {
	r := csv.NewReader(reader)
	imported := 0
	var errors []string

	fileHeaders, err := r.Read()
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

		dto, err := convertFunc(fileHeaders, record)
		if err != nil {
			errors = append(errors, "row "+strconv.Itoa(rowIndex)+": "+err.Error())
			rowIndex++
			continue
		}

		if err := createFunc(ctx, dto); err != nil {
			errors = append(errors, "row "+strconv.Itoa(rowIndex)+": "+err.Error())
		} else {
			imported++
		}

		rowIndex++
	}

	return imported, errors
}

func (uc *Usecase[T, TCreate]) ExportCSV(
	ctx context.Context,
	w io.Writer,
	headers []string,
	toRecord func(T) []string,
) error {
	cw := csv.NewWriter(w)

	if err := cw.Write(headers); err != nil {
		return err
	}

	err := uc.streamRepository.StreamAll(ctx, func(e T) error {
		return cw.Write(toRecord(e))
	})

	cw.Flush()
	return err
}

func (uc *Usecase[T, TCreate]) ExportJSON(ctx context.Context, w io.Writer) error {
	enc := json.NewEncoder(w)

	if _, err := w.Write([]byte("[")); err != nil {
		return err
	}

	first := true

	err := uc.streamRepository.StreamAll(ctx, func(e T) error {
		if !first {
			if _, err := w.Write([]byte(",")); err != nil {
				return err
			}
		}
		first = false

		return enc.Encode(e)
	})
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("]"))
	return err
}
