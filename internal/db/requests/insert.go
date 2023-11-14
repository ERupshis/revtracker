package requests

import (
	"context"
	"fmt"

	"gopkg.in/reform.v1"
)

func InsertOrUpdate(ctx context.Context, db *reform.DB, tx *reform.TX, record reform.Record) error {
	var err error
	if tx != nil {
		err = tx.Save(record)
	} else {
		err = db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
			return tx.Save(record)
		})
	}
	if err != nil {
		return fmt.Errorf("insert/update %s: %w", record.Table().Name(), err)
	}

	return nil
}
