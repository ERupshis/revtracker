package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertQuestion(ctx context.Context, question *data.Question) error {
	return r.insertOrUpdateQuestion(ctx, nil, question)
}

func (r *Reform) UpdateQuestion(ctx context.Context, question *data.Question) error {
	return r.insertOrUpdateQuestion(ctx, nil, question)
}

func (r *Reform) SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error) {
	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	content, err := r.db.WithContext(ctx).SelectOneFrom(data.UserTable, tail, values...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select question by ID: %w", err)
	}

	if content == nil {
		return nil, nil
	}

	return content.(*data.Question), nil
}

func (r *Reform) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return r.deleteQuestionByID(ctx, nil, ID)
}

func (r *Reform) insertOrUpdateQuestion(ctx context.Context, tx *reform.TX, question *data.Question) error {
	var err error
	if tx != nil {
		err = tx.Save(question)
	} else {
		err = r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
			return tx.Save(question)
		})
	}

	if err != nil {
		return fmt.Errorf("failed insert/update question: %w", err)
	}

	return nil
}

func (r *Reform) deleteQuestionByID(ctx context.Context, tx *reform.TX, ID int64) error {
	deleteFunc := func(tx *reform.TX) error {
		tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
		deletedCount, err := tx.DeleteFrom(data.ContentTable, tail, values...)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("delete question by ID: %w", err)
		}

		if deletedCount != 1 {
			_ = tx.Rollback()
			return fmt.Errorf("delete question by ID wrong deletions count: %d", deletedCount)
		}

		return nil
	}

	if tx != nil {
		return deleteFunc(tx)
	} else {
		return r.db.InTransactionContext(ctx, nil, deleteFunc)
	}
}
