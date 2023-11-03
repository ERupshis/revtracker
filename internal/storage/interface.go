package storage

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

//go:generate mockgen -destination=../../mocks/mock_BaseStorage.go -package=mocks github.com/erupshis/revtracker/internal/storage BaseStorage
type BaseStorage interface {
	InsertUser(ctx context.Context, user *data.User) (int64, error)
	SelectUser(ctx context.Context, filters map[string]interface{}) (*data.User, error)

	InsertHomework(ctx context.Context, homework *data.Homework) error
	UpdateHomeworkByID(ctx context.Context, ID int64, newName string) error
	SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error)
	DeleteHomeworkByID(ctx context.Context, ID int64) error

	InsertContent(ctx context.Context, content *data.Content) error
	UpdateContentByID(ctx context.Context, ID int64, values map[string]interface{}) error
	SelectContentByID(ctx context.Context, ID int64) (*data.Content, error)
	DeleteContentByID(ctx context.Context, ID int64) error

	InsertQuestion(ctx context.Context, question *data.Question) error
	UpdateQuestionByID(ctx context.Context, ID int64, values map[string]interface{}) error
	SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error)
	DeleteQuestionByID(ctx context.Context, ID int64) error

	InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error
	UpdateHomeworkQuestionByID(ctx context.Context, ID int64, values map[string]interface{}) error
	SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error)
	DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error
}
