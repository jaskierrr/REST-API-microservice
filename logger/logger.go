package logger

import (
	"card-project/logger/prettylog"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	opts := prettylog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	log := slog.New(handler)
	slog.SetDefault(log)

	return log
}
