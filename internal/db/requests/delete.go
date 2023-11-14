package requests

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func Delete(ctx context.Context, db *reform.DB, tx *reform.TX, filters map[string]interface{}, table reform.Table) error {
	deleteFunc := func(tx *reform.TX) error {
		tail, values := utils.CreateTailAndParams(db, filters)
		_, err := tx.DeleteFrom(table, tail, values...)
		if err != nil {
			return fmt.Errorf("delete %s by filters '%v': %w", table.Name(), filters, err)
		}

		return nil
	}

	if tx != nil {
		return deleteFunc(tx)
	} else {
		return db.InTransactionContext(ctx, nil, deleteFunc)
	}
}
