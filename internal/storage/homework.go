package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertHomework(ctx context.Context, homework *data.Homework) error {
	return s.mngr.InsertHomework(ctx, homework)
}

func (s *Storage) UpdateHomeworkNameByID(ctx context.Context, ID int64, newName string) error {
	return s.mngr.UpdateHomeworkNameByID(ctx, ID, newName)
}

func (s *Storage) SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error) {
	return s.mngr.SelectHomeworkByID(ctx, ID)
}

func (s *Storage) DeleteHomeworkByID(ctx context.Context, ID int64) error {
	return s.mngr.DeleteHomeworkByID(ctx, ID)
}
