package reform

import (
	"context"
	"fmt"
	"sort"

	"github.com/erupshis/revtracker/internal/data"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/db/requests"
	"github.com/erupshis/revtracker/internal/db/utils"
	"gopkg.in/reform.v1"
)

func (r *Reform) InsertData(ctx context.Context, inData *data.Data) error {
	return r.insertOrUpdateData(ctx, inData)
}

func (r *Reform) UpdateData(ctx context.Context, inData *data.Data) error {
	return r.insertOrUpdateData(ctx, inData)
}

func (r *Reform) SelectDataAll(ctx context.Context) ([]data.Data, error) {
	var res []data.Data
	err := r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		homeworks, err := r.selectHomeworks(ctx, tx, nil)
		if err != nil {
			return fmt.Errorf("select all homewrokData: %w", err)
		}

		for _, homework := range homeworks {
			homeworkData, err := r.selectDataHomeworkByID(ctx, tx, homework.ID)
			if err != nil {
				return fmt.Errorf("select all homewrokData: %w", err)
			}

			if homeworkData != nil {
				res = append(res,
					data.Data{
						Homework: *homeworkData,
					},
				)
			}
		}

		return nil
	})

	return res, err
}

func (r *Reform) SelectDataByHomeworkID(ctx context.Context, ID int64) (*data.Data, error) {
	var res *data.Data
	err := r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		homeworkData, err := r.selectDataHomeworkByID(ctx, tx, ID)
		if err != nil {
			return fmt.Errorf("select homeworkData: %w", err)
		}

		if homeworkData != nil {
			res = &data.Data{
				Homework: *homeworkData,
			}
		}

		return nil
	})

	return res, err
}

func (r *Reform) DeleteDataByHomeworkID(ctx context.Context, ID int64) error {
	return r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		if err := requests.Delete(ctx, r.db, tx, []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, ID)}, data.HomeworkQuestionTable); err != nil {
			return fmt.Errorf("delete links in homework_question: %w", err)
		}

		if err := requests.Delete(ctx, r.db, tx, []utils.Argument{utils.CreateArgument(constants.ColID, ID)}, data.HomeworkTable); err != nil {
			return fmt.Errorf("delete homework: %w", err)
		}

		return nil
	})
}

func (r *Reform) insertOrUpdateData(ctx context.Context, inData *data.Data) error {
	homework := &data.Homework{
		ID:   inData.Homework.ID,
		Name: inData.Homework.Name,
	}
	questions := inData.Homework.Questions

	return r.db.InTransactionContext(ctx, nil, func(tx *reform.TX) error {
		if err := requests.InsertOrUpdate(ctx, r.db, tx, homework); err != nil {
			return fmt.Errorf("insert/update homework: %w", err)
		}

		if err := requests.Delete(ctx, r.db, tx, []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, homework.ID)}, data.HomeworkQuestionTable); err != nil {
			return fmt.Errorf("delete links in homework_question: %w", err)
		}

		if err := r.insertQuestions(ctx, tx, questions, homework.ID); err != nil {
			return fmt.Errorf("insert/update questions: %w", err)
		}

		inData.Homework.ID = homework.ID
		inData.Homework.Name = homework.Name

		return nil
	})
}

// TODO: need to look for existing elements and ignore creation new ones.
func (r *Reform) insertQuestions(ctx context.Context, tx *reform.TX, questions []data.Question, homeworkID int64) error {
	for i := 0; i < len(questions); i++ {
		question := &questions[i]

		if err := requests.InsertOrUpdate(ctx, r.db, tx, &question.Content); err != nil {
			return fmt.Errorf("insert/update content: %w", err)
		}

		question.ContentID = question.Content.ID
		if err := requests.InsertOrUpdate(ctx, r.db, tx, question); err != nil {
			return fmt.Errorf("insert/update question: %w", err)
		}

		homeworkQuestion := &data.HomeworkQuestion{
			HomeworkID: homeworkID,
			QuestionID: question.ID,
			Order:      int64(i),
		}

		if err := requests.InsertOrUpdate(ctx, r.db, tx, homeworkQuestion); err != nil {
			return fmt.Errorf("insert/update homework-question link. element's order '%d'(%w)", i, err)
		}
	}
	return nil
}

func (r *Reform) getOrderedQuestionIDs(ctx context.Context, tx *reform.TX, homeworkID int64) ([]int64, error) {
	homeworkQuestions, err := r.selectHomeworkQuestions(ctx, tx, []utils.Argument{utils.CreateArgument(constants.ColHomeworkID, homeworkID)})
	if err != nil {
		return nil, fmt.Errorf("select questions: %w", err)
	}

	sort.Slice(homeworkQuestions, func(l, r int) bool {
		return homeworkQuestions[l].Order < homeworkQuestions[r].Order
	})

	var res []int64
	for _, hq := range homeworkQuestions {
		res = append(res, hq.QuestionID)
	}

	return res, nil
}

func (r *Reform) getQuestions(ctx context.Context, tx *reform.TX, homeworkID int64) ([]data.Question, error) {
	questionsOrder, err := r.getOrderedQuestionIDs(ctx, tx, homeworkID)
	if err != nil {
		return nil, fmt.Errorf("get ordered questionIDs: %w", err)
	}

	var res []data.Question
	for _, questionID := range questionsOrder {
		question, err := r.selectQuestion(ctx, tx, []utils.Argument{utils.CreateArgument(constants.ColID, questionID)})
		if err != nil {
			return nil, fmt.Errorf("get question from db: %w", err)
		}

		res = append(res, *question)
	}

	return res, nil
}

func (r *Reform) selectDataHomeworkByID(ctx context.Context, tx *reform.TX, ID int64) (*data.HomeworkData, error) {
	homework, err := r.selectHomework(ctx, tx, []utils.Argument{utils.CreateArgument(constants.ColID, ID)})
	if err != nil {
		return nil, fmt.Errorf("select homework: %w", err)
	}

	if homework == nil {
		return nil, nil
	}

	questions, err := r.getQuestions(ctx, tx, homework.ID)
	if err != nil {
		return nil, fmt.Errorf("select questions: %w", err)
	}

	return &data.HomeworkData{
		ID:        homework.ID,
		Name:      homework.Name,
		Questions: questions,
	}, nil
}
