package homework

import (
	"errors"
	"fmt"

	"github.com/erupshis/revtracker/internal/controller/handlers/utils"
	"github.com/erupshis/revtracker/internal/db/constants"
	"github.com/erupshis/revtracker/internal/logger"
	"github.com/erupshis/revtracker/internal/storage"
	storageErrors "github.com/erupshis/revtracker/internal/storage/errors"
	"github.com/gofiber/fiber/v2"
)

// Delete func.
// @Description Delete godoc
// @Tags homework
// @Summary deletes existing homework
// @ID homework-delete
// @Success 200
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 500
// @Param id path int true "homework id"
// @Router /homework/{id} [delete]
// @Security ApiKeyAuth
func Delete(storage storage.BaseStorage, log logger.BaseLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := utils.GetIDFromParams(c)
		if err != nil {
			log.Info("%s get ID from params: %v", fmt.Sprintf(packagePath, constants.Select), err)
			c.Status(fiber.StatusBadRequest)
			return nil
		}

		if err = storage.DeleteHomeworkByID(c.Context(), ID); err != nil {
			if errors.Is(err, storageErrors.ErrNoContent) {
				c.Status(fiber.StatusNoContent)
			} else {
				c.Status(fiber.StatusInternalServerError)
			}
			log.Info("%s failed to delete: %v", fmt.Sprintf(packagePath, constants.Delete), err)
			return nil
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
