package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

func (s *Storage) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return nil
}

func (s *Storage) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return nil
}

func (s *Storage) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	return nil, nil
}

func (s *Storage) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return nil
}
