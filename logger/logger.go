package logger

import (
	"log/slog"
	"os"
)

var jsonHandler = slog.NewJSONHandler(os.Stderr, nil)
var logger = slog.New(jsonHandler)

func GetLogger() *slog.Logger {
	return logger
}
