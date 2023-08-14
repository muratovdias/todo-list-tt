package logger

import (
	"golang.org/x/exp/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	var log *slog.Logger
	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return log
}
