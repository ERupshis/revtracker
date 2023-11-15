package requests

import (
	"context"
	"fmt"

	"gopkg.in/reform.v1"
)

// TODO: rename on Insert.
func InsertOrUpdate(ctx context.Context, db *reform.DB, tx *reform.TX, record reform.Record) error {
	insertOrUpdateFunc := func(tx *reform.TX) error {
		return tx.Save(record)
	}

	var err error
	if tx != nil {
		err = insertOrUpdateFunc(tx)
	} else {
		err = db.InTransactionContext(ctx, nil, insertOrUpdateFunc)
	}

	if err != nil {
		return fmt.Errorf("insert %s: %w", record.Table().Name(), err)
	}

	return nil
}
