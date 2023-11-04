package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertData(ctx context.Context, data *data.Data) error {

	return nil
}

func (s *Storage) UpdateData(ctx context.Context, data *data.Data) error {

	return nil
}

func (s *Storage) SelectDataByHomeworkID(ctx context.Context) (*data.Data, error) {

	return nil, nil
}

func (s *Storage) DeleteDataByHomeworkID(ctx context.Context, ID int64) error {

	return nil
}
