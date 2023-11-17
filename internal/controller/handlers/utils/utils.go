package utils

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

var ErrMissingIDinURI = fmt.Errorf("missing ID in URI")

func GetIDFromParams(c *fiber.Ctx) (int64, error) {
	if c == nil {
		return -1, fmt.Errorf("fiber ctx is nil")
	}

	rawID := c.Params("ID")
	if rawID == "" {
		return -1, ErrMissingIDinURI
	}

	ID, err := strconv.Atoi(rawID)
	if err != nil {
		return -1, fmt.Errorf("parse ID from URI: %w", err)
	}

	if ID <= 0 {
		return -1, fmt.Errorf("ID is below ir equal 0")
	}

	return int64(ID), nil
}

func IsUniqueConstraint(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return true
		}
	}

	return false
}

func IsForeignKeyConstraint(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23503" {
			return true
		}
	}

	return false
}
