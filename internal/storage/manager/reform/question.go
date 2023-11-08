package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertQuestion(ctx context.Context, question *data.Question) error {
	return r.insertQuestionAndContent(ctx, nil, question)
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

func (r *Reform) insertQuestionAndContent(ctx context.Context, tx *reform.TX, question *data.Question) error {
	insertOrUpdateFunc := func(tx *reform.TX) error {
		if err := common.InsertOrUpdate(ctx, r.db, tx, &question.Content); err != nil {
			return fmt.Errorf("inser question: add content: %w", err)
		}

		question.ContentID = question.Content.ID

		if err := common.InsertOrUpdate(ctx, r.db, tx, question); err != nil {
			return fmt.Errorf("inser question: add question: %w", err)
		}

		return nil
	}

	if tx != nil {
		return insertOrUpdateFunc(tx)
	}

	return r.db.InTransactionContext(ctx, nil, insertOrUpdateFunc)
}

func (r *Reform) selectQuestion(ctx context.Context, tx *reform.TX, filters map[string]interface{}) (*data.Question, error) {
	var question *data.Question

	selectFunc := func(tx *reform.TX) error {
		questionRaw, err := common.SelectOne(ctx, r.db, tx, filters, data.QuestionTable)
		if err != nil {
			return fmt.Errorf("select question by filters '%v': %w", filters, err)
		}

		question = questionRaw.(*data.Question)
		questionContent, err := r.selectContent(ctx, tx, map[string]interface{}{"id": question.ContentID})
		if err != nil {
			return fmt.Errorf("select question by id '%d': %w", question.ContentID, err)
		}

		question.Content = *questionContent

		return nil
	}

	if tx != nil {
		err := selectFunc(tx)
		return question, err
	}

	err := r.db.InTransactionContext(ctx, nil, selectFunc)
	return question, err
}
