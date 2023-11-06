package utils

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetIDFromParams(c *fiber.Ctx) (int64, error) {
	if c == nil {
		return -1, fmt.Errorf("fiber ctx is nil")
	}

	rawID := c.Params("ID")
	if rawID == "" {
		return -1, fmt.Errorf("missing ID in URI")
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
