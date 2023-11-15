package requests

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func Update(ctx context.Context, db *reform.DB, tx *reform.TX, record reform.Record) error {
	tail, values := utils.CreateTailAndParams(db, []utils.Argument{utils.CreateArgument(constants.ColID, record.PKValue())}, 0)

	insertOrUpdateFunc := func(tx *reform.TX) error {
		reformStruct, err := tx.SelectOneFrom(record.View(), utils.AddDeletedCheck(tail, false), values...)
		if err != nil {
			return fmt.Errorf("try to select elem: %w", err)
		}

		updateExistingRecord(record, reformStruct)
		return tx.Update(record)
	}

	var err error
	if tx != nil {
		err = insertOrUpdateFunc(tx)
	} else {
		err = db.InTransactionContext(ctx, nil, insertOrUpdateFunc)
	}

	if err != nil {
		return fmt.Errorf("update %s: %w", record.Table().Name(), err)
	}

	return nil
}

func updateExistingRecord(record reform.Record, reformStruct reform.Struct) {
	switch rec := record.(type) {
	case *data.Question:
		existingRecord := reformStruct.(*data.Question)
		rec.ContentID = existingRecord.ContentID
	default:
		return
	}
}
