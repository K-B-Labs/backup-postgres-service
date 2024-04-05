package logger

import (
	"log/slog"
	"os"
	"sync"
)

type jsonLogger struct {
	instance *slog.Logger
}

var (
	logger *jsonLogger
	once   sync.Once
)

func GetLogger() *slog.Logger {
	once.Do(func() {
		logger = &jsonLogger{
			instance: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
		}
	})

	return logger.instance
}
