package logger

import (
	"log/slog"
)

const (
	Warning  = 1
	Error    = 2
	Previous = "previous_error"
)

func LogErrors(severity int, err error, logger *slog.Logger, items map[string]interface{}) {
	if severity == Warning {
		logger.Warn(err.Error())
	} else {
		logger.Error(err.Error())
	}
}

func Log(err error, logger *slog.Logger) {
	items := make(map[string]interface{}, 0)
	LogWithItems(err, logger, items)
}

func LogWithItems(err error, logger *slog.Logger, extraItems map[string]interface{}) {
	severity := Error

	LogErrors(severity, err, logger, extraItems)
}
