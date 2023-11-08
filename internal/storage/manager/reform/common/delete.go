package common

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/storage/manager/reform/utils"
	"gopkg.in/reform.v1"
)

func Delete(ctx context.Context, db *reform.DB, tx *reform.TX, filters map[string]interface{}, table reform.Table) error {
	deleteFunc := func(tx *reform.TX) error {
		tail, values := utils.CreateTailAndParams(db, filters)
		deletedCount, err := tx.DeleteFrom(table, tail, values...)
		if err != nil {
			return fmt.Errorf("delete %s by ID: %w", table.Name(), err)
		}

		if deletedCount > 1 {
			return fmt.Errorf("delete %s by ID wrong deletions count: %d", table.Name(), deletedCount)
		}

		return nil
	}

	if tx != nil {
		return deleteFunc(tx)
	} else {
		return db.InTransactionContext(ctx, nil, deleteFunc)
	}
}
