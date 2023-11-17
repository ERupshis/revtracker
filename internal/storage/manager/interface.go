package manager

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
)

type BaseStorageManager interface {
	InsertData(ctx context.Context, data *data.Data) error
	UpdateData(ctx context.Context, data *data.Data) error
	SelectDataAll(ctx context.Context) ([]data.Data, error)
	SelectDataByHomeworkID(ctx context.Context, ID int64) (*data.Data, error)
	DeleteDataByHomeworkID(ctx context.Context, ID int64) error

	InsertHomework(ctx context.Context, homework *data.Homework) error
	UpdateHomework(ctx context.Context, homework *data.Homework) error
	SelectHomeworks(ctx context.Context) ([]data.Homework, error)
	SelectHomeworkByID(ctx context.Context, ID int64) (*data.Homework, error)
	DeleteHomeworkByID(ctx context.Context, ID int64) error

	InsertContent(ctx context.Context, content *data.Content) error
	UpdateContent(ctx context.Context, content *data.Content) error
	SelectContentByID(ctx context.Context, ID int64) (*data.Content, error)
	DeleteContentByID(ctx context.Context, ID int64) error

	InsertQuestion(ctx context.Context, question *data.Question) error
	UpdateQuestion(ctx context.Context, question *data.Question) error
	SelectQuestions(ctx context.Context) ([]data.Question, error)
	SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error)
	DeleteQuestionByID(ctx context.Context, ID int64) error

	InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error
	UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error
	SelectHomeworkQuestions(ctx context.Context) ([]data.HomeworkQuestion, error)
	SelectHomeworkQuestionsByHomeworkID(ctx context.Context, ID int64) ([]data.HomeworkQuestion, error)
	SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error)
	DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error
}
