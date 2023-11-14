package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	common2 "github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return common2.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) UpdateHomeworkQuestion(ctx context.Context, homeworkQuestion *data.HomeworkQuestion) error {
	return common2.InsertOrUpdate(ctx, r.db, nil, homeworkQuestion)
}

func (r *Reform) SelectHomeworkQuestions(ctx context.Context) ([]data.HomeworkQuestion, error) {
	return r.selectHomeworkQuestions(ctx, nil, nil)
}

func (r *Reform) SelectHomeworkQuestionsByHomeworkID(ctx context.Context, ID int64) ([]data.HomeworkQuestion, error) {
	return r.selectHomeworkQuestions(ctx, nil, []utils.Argument{utils.CreateArgument("homework_id", ID)})
}

func (r *Reform) SelectHomeworkQuestionByID(ctx context.Context, ID int64) (*data.HomeworkQuestion, error) {
	content, err := common2.SelectOne(ctx, r.db, nil, []utils.Argument{utils.CreateArgument("id", ID)}, data.HomeworkQuestionTable)
	return content.(*data.HomeworkQuestion), err
}

func (r *Reform) DeleteHomeworkQuestionByID(ctx context.Context, ID int64) error {
	return common2.Delete(ctx, r.db, nil, []utils.Argument{utils.CreateArgument("id", ID)}, data.HomeworkQuestionTable)
}

func (r *Reform) selectHomeworkQuestions(ctx context.Context, tx *reform.TX, filters []utils.Argument) ([]data.HomeworkQuestion, error) {
	content, err := common2.SelectAll(ctx, r.db, tx, filters, " homework_id ASC, \"order\" ASC, id ASC", data.HomeworkQuestionTable)
	if err != nil {
		return nil, fmt.Errorf("select homeworkQuestions: %w", err)
	}

	var res []data.HomeworkQuestion
	for _, elem := range content {
		res = append(res, *(elem.(*data.HomeworkQuestion)))
	}
	return res, nil
}
