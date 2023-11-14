package requests

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	authData "github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func InsertOrUpdate(ctx context.Context, db *reform.DB, tx *reform.TX, record reform.Record) error {
	uniqueFilters := getUniqueFilters(record)
	tail, values := utils.CreateTailAndParams(db, uniqueFilters, 0)

	insertOrUpdateFunc := func(tx *reform.TX) error {
		reformStruct, err := tx.SelectOneFrom(record.View(), utils.AddDeletedCheck(tail, true), values...)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("try to select deleted elem: %w", err)
		}

		if reformStruct != nil && len(uniqueFilters) != 0 {
			record.SetPK(reformStruct.(reform.Record).PKValue())
		}

		return tx.Save(record)
	}

	var err error
	if tx != nil {
		err = insertOrUpdateFunc(tx)
	} else {
		err = db.InTransactionContext(ctx, nil, insertOrUpdateFunc)
	}

	if err != nil {
		return fmt.Errorf("insert/update %s: %w", record.Table().Name(), err)
	}

	return nil
}

func getUniqueFilters(record reform.Record) []utils.Argument {
	switch rec := record.(type) {
	case *data.Homework:
		return []utils.Argument{utils.CreateArgument(constants.ColName, rec.Name)}
	case *data.HomeworkQuestion:
		return []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, rec.HomeworkID), utils.CreateArgument(constants.ColQuestionID, rec.QuestionID)}
	case *data.Question:
		return []utils.Argument{utils.CreateArgument(constants.ColName, rec.Name)}
	case *authData.User:
		return []utils.Argument{utils.CreateArgument(constants.ColName, rec.Name), utils.CreateArgumentAND(constants.ColLogin, rec.Login)}
	default:
		panic("unknown type")
	}

	return nil
}
