package utils

import (
	"context"
	"fmt"
	"strconv"

	"github.com/erupshis/revtracker/internal/auth/data"
	"github.com/erupshis/revtracker/internal/auth/middleware"
)

const wrongUserDataMsgTemplate = "incorrect user data:"

func GetUserIDFromContext(ctx context.Context) (int64, error) {
	userIDraw := ctx.Value(middleware.ContextString(data.UserID))
	if userIDraw == nil {
		return -1, fmt.Errorf("missing userID in request's context")
	}

	userID, err := strconv.ParseInt(userIDraw.(string), 10, 64)
	if err != nil {
		return -1, fmt.Errorf("parse userID from request's context: %w", err)
	}

	return userID, nil
}

func IsUserDataValid(userData *data.User) (bool, error) {
	if userData == nil {
		return false, fmt.Errorf("%s userData is nil", wrongUserDataMsgTemplate)
	}

	var errMsg string
	if userData.Login == "" {
		errMsg += " login"
	}

	if userData.Password == "" {
		errMsg += " password"
	}

	if userData.Name == "" {
		errMsg += " name"
	}

	if errMsg != "" {
		return false, fmt.Errorf("%s%s", wrongUserDataMsgTemplate, errMsg)
	}

	return true, nil
}
