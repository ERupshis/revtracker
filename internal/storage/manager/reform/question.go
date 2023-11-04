package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
)

func (r *Reform) InsertQuestion(ctx context.Context, question *data.Question) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for question insert: %w", err)
	}

	err = tx.Save(question)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update question name by ID: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) UpdateQuestion(ctx context.Context, question *data.Question) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for question update: %w", err)
	}

	err = tx.Update(question)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update question: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
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
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for question delete: %w", err)
	}

	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	deletedCount, err := tx.DeleteFrom(data.ContentTable, tail, values...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete question by ID: %w", err)
	}

	if deletedCount != 1 {
		_ = tx.Rollback()
		return fmt.Errorf("delete content by ID wrong deletions count: %d", deletedCount)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}
