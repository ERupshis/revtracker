package logger

import (
	"github.com/gofiber/fiber/v2"
)

type testPlug struct {
}

func CreateTestPLug() (BaseLogger, error) {
	return &testPlug{}, nil
}

func (t *testPlug) Info(msg string, fields ...interface{}) {
}

func (t *testPlug) Printf(msg string, fields ...interface{}) {
}

func (t *testPlug) Sync() {
}

func (t *testPlug) LogHandler(c *fiber.Ctx) error {
	return nil
}
