package manager

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

type BaseStorageManager interface {
	InsertUser(ctx context.Context, user *data.User) (int64, error)
	SelectUser(ctx context.Context, filters map[string]interface{}) (*data.User, error)

	InsertHomework(ctx context.Context, homework *data.Homework) error
	UpdateHomework(ctx context.Context, homework *data.Homework) error
	SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error)
	DeleteHomeworkByID(ctx context.Context, ID int64) error

	InsertContent(ctx context.Context, content *data.Content) error
	UpdateContent(ctx context.Context, content *data.Content) error
	SelectContentByID(ctx context.Context, ID int64) (*data.Content, error)
	DeleteContentByID(ctx context.Context, ID int64) error

	InsertQuestion(ctx context.Context, question *data.Question) error
	UpdateQuestion(ctx context.Context, question *data.Question) error
	SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error)
	DeleteQuestionByID(ctx context.Context, ID int64) error
}
