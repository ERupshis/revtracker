package reform

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
)

func (r *Reform) InsertContent(ctx context.Context, content *data.Content) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for content update: %w", err)
	}

	err = tx.Save(content)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update homework name by ID: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) UpdateContent(ctx context.Context, content *data.Content) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for content update: %w", err)
	}

	err = tx.Update(content)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update homework name by ID: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) SelectContentByID(ctx context.Context, ID int64) (*data.Content, error) {
	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	content, err := r.db.WithContext(ctx).SelectOneFrom(data.UserTable, tail, values...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("select content by ID: %w", err)
	}

	if content == nil {
		return nil, nil
	}

	return content.(*data.Content), nil
}

func (r *Reform) DeleteContentByID(ctx context.Context, ID int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction for content delete: %w", err)
	}

	tail, values := utils.CreateTailAndParams(r.db, map[string]interface{}{"id": ID})
	deletedCount, err := tx.DeleteFrom(data.ContentTable, tail, values...)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete content by ID: %w", err)
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
