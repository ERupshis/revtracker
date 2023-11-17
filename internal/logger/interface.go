// Package loggerZap provides interface for loggerZap in project.
package logger

import (
	"github.com/gofiber/fiber/v2"
)

// BaseLogger interface of used loggerZap.
type BaseLogger interface {
	// Sync flushing any buffered log entries.
	Sync()

	// Info generates 'info' level log.
	Info(msg string, fields ...interface{})

	// Printf interface for kafka's implementation.
	Printf(msg string, fields ...interface{})

	// LogHandler handler for requests logging.
	LogHandler(c *fiber.Ctx) error
}
