package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertQuestion(ctx context.Context, question *data.Question) error {
	return nil
}

func (s *Storage) UpdateQuestion(ctx context.Context, question *data.Question) error {
	return nil
}

func (s *Storage) SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error) {
	return nil, nil
}

func (s *Storage) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return nil
}
