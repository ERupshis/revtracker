package reform

import (
	"context"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertQuestion(ctx context.Context, question *data.Question) error {
	return common.InsertOrUpdate(ctx, r.db, nil, question)
}

func (r *Reform) UpdateQuestion(ctx context.Context, question *data.Question) error {
	return common.InsertOrUpdate(ctx, r.db, nil, question)
}

func (r *Reform) SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error) {
	return r.selectQuestion(ctx, nil, map[string]interface{}{"id": ID})
}

func (r *Reform) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return common.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.QuestionTable)
}

func (r *Reform) selectQuestion(ctx context.Context, tx *reform.TX, filters map[string]interface{}) (*data.Question, error) {
	content, err := common.SelectOne(ctx, r.db, tx, filters, data.QuestionTable)
	return content.(*data.Question), err
}
