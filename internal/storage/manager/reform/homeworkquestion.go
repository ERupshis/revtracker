package reform

import (
	"context"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	storageErrors "github.com/erupshis/revtracker/internal/storage/errors"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	homeworkQuestions, err := r.selectHomeworkQuestions(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, homeworkQuestion.HomeworkID)})
	if err != nil && !errors.Is(err, storageErrors.ErrNoContent) {
		return fmt.Errorf("select homework questions by ID '%d'", homeworkQuestion.ID)
	}

	var maxOrderNum int64 = 0
	for _, hq := range homeworkQuestions {
		if maxOrderNum < hq.Order {
			maxOrderNum = hq.Order
		}
	}

	homeworkQuestion.Order = maxOrderNum
	if len(homeworkQuestions) != 0 {
		homeworkQuestion.Order++
	}
	return requests.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	currHomeworkQuestion, err := r.selectHomeworkQuestion(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColID, homeworkQuestion.ID)})
	if err != nil {
		return fmt.Errorf("select existing record: %w", err)
	}

	homeworkQuestion.Order = currHomeworkQuestion.Order
	homeworkQuestion.HomeworkID = currHomeworkQuestion.HomeworkID
	return requests.Update(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) SelectHomeworkQuestions(ctx context.Context) ([]data.HomeworkQuestion, error) {
	return r.selectHomeworkQuestions(ctx, nil, nil)
}

func (r *Reform) SelectHomeworkQuestionsByHomeworkID(ctx context.Context, ID int64) ([]data.HomeworkQuestion, error) {
	return r.selectHomeworkQuestions(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, ID)})
}

func (r *Reform) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	return r.selectHomeworkQuestion(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)})
}

func (r *Reform) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)}, data.HomeworkQuestionTable)
}

func (r *Reform) selectHomeworkQuestion(ctx context.Context, tx *reform.TX, filters []utils.Argument) (*data.HomeworkQuestion, error) {
	content, err := requests.SelectOne(ctx, r.db, tx, filters, data.HomeworkQuestionTable)

	if content == nil {
		return nil, err
	}

	return content.(*data.HomeworkQuestion), err
}

func (r *Reform) selectHomeworkQuestions(ctx context.Context, tx *reform.TX, filters []utils.Argument) ([]data.HomeworkQuestion, error) {
	ordering := fmt.Sprintf(" %s ASC, %s ASC, %s ASC", constants.ColHomeworkID, constants.ColOrder, constants.ColID)
	content, err := requests.SelectAll(ctx, r.db, tx, filters, ordering, data.HomeworkQuestionTable)
	if err != nil {
		return nil, fmt.Errorf("select homeworkQuestions: %w", err)
	}

	var res []data.HomeworkQuestion
	for _, elem := range content {
		res = append(res, *(elem.(*data.HomeworkQuestion)))
	}
	return res, nil
}
