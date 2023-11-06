package logger

import (
	"github.com/gofiber/fiber/v2"
)

type logMock struct {
}

func CreateMock() (BaseLogger, error) {
	return &logMock{}, nil
}

func (t *logMock) Info(_ string, _ ...interface{}) {
}

func (t *logMock) Printf(_ string, _ ...interface{}) {
}

func (t *logMock) Sync() {
}

func (t *logMock) LogHandler(_ *fiber.Ctx) error {
	return nil
}
