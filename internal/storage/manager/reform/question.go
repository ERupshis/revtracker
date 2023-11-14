package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/requests"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertQuestion(ctx context.Context, question *data.Question) error {
	return r.insertQuestionAndContent(ctx, nil, question)
}

func (r *Reform) UpdateQuestion(ctx context.Context, question *data.Question) error {
	return r.insertQuestionAndContent(ctx, nil, question)
}

func (r *Reform) SelectQuestions(ctx context.Context) ([]data.Question, error) {
	return r.selectQuestions(ctx, nil, nil)
}

func (r *Reform) SelectQuestionByID(ctx context.Context, ID int64) (*data.Question, error) {
	return r.selectQuestion(ctx, nil, map[string]interface{}{"id": ID})
}

func (r *Reform) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, map[string]interface{}{"id": ID}, data.QuestionTable)
}

func (r *Reform) insertQuestionAndContent(ctx context.Context, tx *reform.TX, question *data.Question) error {
	insertOrUpdateFunc := func(tx *reform.TX) error {
		currentQuestion, err := r.selectQuestion(ctx, nil, map[string]interface{}{"id": question.ID})
		if err != nil {
			return fmt.Errorf("insert question: check question in db: %w", err)
		}

		if currentQuestion != nil {
			question.ContentID = currentQuestion.ContentID
			question.Content.ID = currentQuestion.ContentID
		}

		if err = requests.InsertOrUpdate(ctx, r.db, tx, &question.Content); err != nil {
			return fmt.Errorf("insert question: add content: %w", err)
		}

		question.ContentID = question.Content.ID

		if err = requests.InsertOrUpdate(ctx, r.db, tx, question); err != nil {
			return fmt.Errorf("insert question: add question: %w", err)
		}

		return nil
	}

	if tx != nil {
		return insertOrUpdateFunc(tx)
	}

	return r.db.InTransactionContext(ctx, nil, insertOrUpdateFunc)
}

// TODO: need to add custom query.
func (r *Reform) selectQuestions(ctx context.Context, tx *reform.TX, filters map[string]interface{}) ([]data.Question, error) {
	var questions []data.Question

	selectFunc := func(tx *reform.TX) error {
		questionsRaw, err := requests.SelectAll(ctx, r.db, tx, filters, "id", data.QuestionTable)
		if err != nil {
			return fmt.Errorf("select question by filters '%v': %w", filters, err)
		}

		if questionsRaw == nil {
			return nil
		}

		for _, q := range questionsRaw {
			questions = append(questions, *q.(*data.Question))

			questionContent, err := r.selectContent(ctx, tx, map[string]interface{}{"id": questions[len(questions)-1].ContentID})
			if err != nil {
				return fmt.Errorf("select question by id '%d': %w", questions[len(questions)-1].ContentID, err)
			}

			questions[len(questions)-1].Content = *questionContent
		}

		return nil
	}

	if tx != nil {
		err := selectFunc(tx)
		return questions, err
	}

	err := r.db.InTransactionContext(ctx, nil, selectFunc)
	return questions, err
}

func (r *Reform) selectQuestion(ctx context.Context, tx *reform.TX, filters map[string]interface{}) (*data.Question, error) {
	var question *data.Question

	selectFunc := func(tx *reform.TX) error {
		questionRaw, err := requests.SelectOne(ctx, r.db, tx, filters, data.QuestionTable)
		if err != nil {
			return fmt.Errorf("select question by filters '%v': %w", filters, err)
		}

		if questionRaw == nil {
			return nil
		}

		question = questionRaw.(*data.Question)
		questionContent, err := r.selectContent(ctx, tx, map[string]interface{}{"id": question.ContentID})
		if err != nil {
			return fmt.Errorf("select question by id '%d': %w", question.ContentID, err)
		}

		question.Content = *questionContent

		return nil
	}

	if tx != nil {
		err := selectFunc(tx)
		return question, err
	}

	err := r.db.InTransactionContext(ctx, nil, selectFunc)
	return question, err
}
