package logger

import (
	"github.com/gofiber/fiber/v2"
)

type testPlug struct {
}

func CreateTestPLug() (BaseLogger, error) {
	return &testPlug{}, nil
}

func (_ *testPlug) Info(msg string, fields ...interface{}) {
}

func (_ *testPlug) Printf(msg string, fields ...interface{}) {
}

func (_ *testPlug) Sync() {
}

func (_ *testPlug) LogHandler(c *fiber.Ctx) error {
	return nil
}
