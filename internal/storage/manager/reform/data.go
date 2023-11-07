package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/storage/manager/reform/common"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertData(ctx context.Context, inData *data.Data) error {
	return r.insertOrUpdateData(ctx, inData)
}

func (r *Reform) UpdateData(ctx context.Context, inData *data.Data) error {
	return r.insertOrUpdateData(ctx, inData)
}

func (r *Reform) SelectDataByHomeworkID(ctx context.Context) (*data.Data, error) {

	return nil, nil
}

func (r *Reform) DeleteDataByHomeworkID(ctx context.Context, ID int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction to insert inData: %w", err)
	}

	if err = common.Delete(ctx, r.db, tx, map[string]interface{}{"homework_id": ID}, data.HomeworkQuestionTable); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete links in homework_question: %w", err)
	}

	if err = common.Delete(ctx, r.db, tx, map[string]interface{}{"id": ID}, data.HomeworkTable); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete homework: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) insertOrUpdateData(ctx context.Context, inData *data.Data) error {
	homework := &inData.Homework
	questions := inData.Homework.Questions

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("create transaction to insert inData: %w", err)
	}

	if err = common.InsertOrUpdate(ctx, r.db, tx, homework); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("insert/update homework: %w", err)
	}

	if err = common.Delete(ctx, r.db, tx, map[string]interface{}{"homework_id": homework.ID}, data.HomeworkQuestionTable); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("delete links in homework_question: %w", err)
	}

	if err = r.insertQuestions(ctx, tx, questions, homework.ID); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("insert/update questions: %w", err)
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("commit transaction: %w", err)
	}

	return nil
}

func (r *Reform) insertQuestions(ctx context.Context, tx *reform.TX, questions []data.Question, homeworkID int64) error {
	for i := 0; i < len(questions); i++ {
		question := &questions[i]
		err := common.InsertOrUpdate(ctx, r.db, tx, &question.Content)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("insert/update content")
		}

		question.ContentID = question.Content.ID
		err = common.InsertOrUpdate(ctx, r.db, tx, question)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("insert/update question")
		}

		homeworkQuestion := &data.HomeworkQuestion{
			HomeworkID: homeworkID,
			QuestionID: question.ID,
			Order:      int64(i),
		}
		err = common.InsertOrUpdate(ctx, r.db, tx, homeworkQuestion)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("insert/update homework-question link. element's order: %d", i)
		}
	}
	return nil
}
