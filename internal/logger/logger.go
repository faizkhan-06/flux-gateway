package logger

import (
	"log/slog"
	"os"
)

func Init() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	slog.SetDefault(logger)
}
