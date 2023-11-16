package errors

import (
	"errors"

	"github.com/jackc/pgconn"
)

func IsLinkBetweenDataProblem(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23503" {
			return true
		}
	}

	return false
}

var ErrQuestionNotFound = &pgconn.PgError{
	Message: "question is not found",
}

func IsQuestionNotFound(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Message == ErrQuestionNotFound.Message {
			return true
		}
	}

	return false
}

var ErrQuestionAlreadyInHomework = &pgconn.PgError{
	Message: "same question already has been added in homework",
}

func IsQuestionAlreadyInHomework(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Message == ErrQuestionAlreadyInHomework.Message {
			return true
		}
	}

	return false
}
