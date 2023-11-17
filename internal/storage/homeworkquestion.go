package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return s.mngr.InsertHomeworkQuestion(ctx, homeworkQuestion)
}

func (s *Storage) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return s.mngr.UpdateHomeworkQuestion(ctx, homeworkQuestion)
}

func (s *Storage) SelectHomeworkQuestions(ctx context.Context) ([]data.HomeworkQuestion, error) {
	return s.mngr.SelectHomeworkQuestions(ctx)
}

func (s *Storage) SelectHomeworkQuestionsByHomeworkID(ctx context.Context, ID int64) ([]data.HomeworkQuestion, error) {
	return s.mngr.SelectHomeworkQuestionsByHomeworkID(ctx, ID)
}

func (s *Storage) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	return s.mngr.SelectHomeworkQuestionByID(ctx, ID)
}

func (s *Storage) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return s.mngr.DeleteHomeworkQuestionByID(ctx, ID)
}
