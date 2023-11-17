package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertQuestion(ctx context.Context, question *data.Question) error {
	return s.mngr.InsertQuestion(ctx, question)
}

func (s *Storage) UpdateQuestion(ctx context.Context, question *data.Question) error {
	return s.mngr.UpdateQuestion(ctx, question)
}

func (s *Storage) SelectQuestions(ctx context.Context) ([]data.Question, error) {
	return s.mngr.SelectQuestions(ctx)
}

func (s *Storage) SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error) {
	return s.mngr.SelectQuestionByID(ctx, ID)
}

func (s *Storage) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return s.mngr.DeleteQuestionByID(ctx, ID)
}
