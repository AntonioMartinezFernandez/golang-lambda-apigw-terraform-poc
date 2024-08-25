package logger

import (
	"log/slog"
	"os"
)

func NewLogger(logLevel string) *slog.Logger {
	level := logLev(logLevel)
	opts := &slog.HandlerOptions{Level: level}

	return slog.New(NewPrettyLogHandler(opts).WithGroup("data"))
}

func NewJsonLogger(logLevel string) *slog.Logger {
	level := logLev(logLevel)

	jsonHandler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	)

	return slog.New(jsonHandler)
}

func NewNullLogger() *slog.Logger {
	return &slog.Logger{}
}

func logLev(lvl string) slog.Level {
	var logLevel slog.Level
	switch level := lvl; level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelWarn
	}

	return logLevel
}
