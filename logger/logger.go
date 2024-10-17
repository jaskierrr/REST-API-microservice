package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	// slog.SetDefault(log)
	return log
}
