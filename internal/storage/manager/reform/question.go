package reform

import (
	"context"
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
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
	return r.selectQuestion(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)}, false)
}

func (r *Reform) DeleteQuestionByID(ctx context.Context, ID int64) error {
	return requests.Delete(ctx, r.db, nil, []utils.Argument{utils.CreateArgument(constants.ColID, ID)}, data.QuestionTable)
}

func (r *Reform) insertQuestionAndContent(ctx context.Context, tx *reform.TX, question *data.Question) error {
	insertOrUpdateFunc := func(tx *reform.TX) error {
		currentQuestion, err := r.selectQuestion(ctx, nil, []utils.Argument{utils.CreateArgument(constants.ColID, question.ID)}, true)
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
func (r *Reform) selectQuestions(ctx context.Context, tx *reform.TX, filters []utils.Argument) ([]data.Question, error) {
	var questions []data.Question

	selectFunc := func(tx *reform.TX) error {
		questionsRaw, err := requests.SelectAll(ctx, r.db, tx, filters, constants.ColID, data.QuestionTable)
		if err != nil {
			return fmt.Errorf("select question by filters '%v': %w", filters, err)
		}

		if questionsRaw == nil {
			return nil
		}

		for _, q := range questionsRaw {
			questions = append(questions, *q.(*data.Question))

			questionContent, err := r.selectContent(ctx, tx, []utils.Argument{utils.CreateArgument(constants.ColID, questions[len(questions)-1].ContentID)})
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

// TODO: refactor selectQuestion methods.
func (r *Reform) selectQuestion(ctx context.Context, tx *reform.TX, filters []utils.Argument, abs bool) (*data.Question, error) {
	var question *data.Question

	selectFunc := func(tx *reform.TX) error {
		var questionRaw reform.Struct
		var err error
		if abs {
			questionRaw, err = requests.SelectOneAbs(ctx, r.db, tx, filters, data.QuestionTable)
		} else {
			questionRaw, err = requests.SelectOne(ctx, r.db, tx, filters, data.QuestionTable)
		}

		if err != nil {
			return fmt.Errorf("select question by filters '%v': %w", filters, err)
		}

		if questionRaw == nil {
			return nil
		}

		question = questionRaw.(*data.Question)
		var questionContent *data.Content
		if abs {
			questionContent, err = r.selectContentAbs(ctx, tx, []utils.Argument{utils.CreateArgument(constants.ColID, question.ContentID)})
		} else {
			questionContent, err = r.selectContent(ctx, tx, []utils.Argument{utils.CreateArgument(constants.ColID, question.ContentID)})
		}

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
