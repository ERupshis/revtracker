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

func (r *Reform) InsertContent(ctx context.Context, content *data.Content) error {
	return r.insertOrUpdateContent(ctx, content)
}

func (r *Reform) UpdateContent(ctx context.Context, content *data.Content) error {
	return r.insertOrUpdateContent(ctx, content)
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
	return r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
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

		return nil
	})
}

func (r *Reform) insertOrUpdateContent(ctx context.Context, content *data.Content) error {
	err := r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		return tx.Save(content)
	})

	if err != nil {
		return fmt.Errorf("failed insert/update content: %w", err)
	}

	return nil
}
