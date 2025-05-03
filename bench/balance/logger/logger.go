package logger

import "log/slog"

var logger *slog.Logger

func GetLogger() *slog.Logger {
	return logger
}

func init() {
	logger = slog.Default()
}
