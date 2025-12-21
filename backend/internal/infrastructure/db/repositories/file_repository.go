package repositories

import (
	"context"
)

type StreamRepository[T any] struct {
	BaseRepository
}

func NewStreamRepository[T any](
	baseRepository BaseRepository,
) *StreamRepository[T] {
	return &StreamRepository[T]{baseRepository}
}

func (r *StreamRepository[T]) StreamAll(
	ctx context.Context,
	fn func(e T) error,
) error {
	var entity T
	rows, err := r.db.WithContext(ctx).
		Model(new(T)).
		Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		if err = r.db.ScanRows(rows, &entity); err != nil {
			return err
		}

		if err = fn(entity); err != nil {
			return err
		}
	}

	return nil
}
