package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertHomework(ctx context.Context, homework *data.Homework) error {
	return nil
}

func (s *Storage) UpdateHomeworkByID(ctx context.Context, ID int64, newName string) error {
	return nil
}

func (s *Storage) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	return nil, nil
}

func (s *Storage) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	return nil
}
