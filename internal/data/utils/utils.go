package utils

import (
	"fmt"

	"github.com/erupshis/revtracker/internal/data"
)

const check = "check:"

func ValidateHomeworkData(data *data.Homework) error {
	if data.Name == "" {
		return fmt.Errorf("empty name")
	}

	return nil
}

func ValidateQuestionsData(data []data.Question) error {
	var problemQuestionIdxs []int

	for i, question := range data {
		if question.Name == "" {
			problemQuestionIdxs = append(problemQuestionIdxs, i)
			continue
		}

		if err := ValidateContentData(&question.Content); err != nil {
			problemQuestionIdxs = append(problemQuestionIdxs, i)
			continue
		}
	}

	if len(problemQuestionIdxs) == 0 {
		return nil
	}

	return fmt.Errorf("invalid questions: need to check idx: %v", problemQuestionIdxs)
}

func ValidateContentData(data *data.Content) error {
	errBody := check

	if data.Task == nil || *data.Task == "" {
		errBody += " task"
	}

	if data.Answer == nil || *data.Answer == "" {
		errBody += " answer"
	}

	if data.Solution == nil || *data.Solution == "" {
		errBody += " solution"
	}

	if errBody == check {
		return nil
	}

	return fmt.Errorf("invalid content data: %s", errBody)
}
