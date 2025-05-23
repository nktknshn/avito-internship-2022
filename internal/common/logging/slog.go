package logging

import (
	"log/slog"
	"os"
)

type LoggerSlog struct {
	logger *slog.Logger
}

func NewSlog() *LoggerSlog {
	return &LoggerSlog{
		logger: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
	}
}

// InitLogger инициализирует логгер
func (l *LoggerSlog) InitLogger(_ ...interface{}) {
	l.logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
}

func (l *LoggerSlog) Debug(msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l *LoggerSlog) Info(msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *LoggerSlog) Warn(msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *LoggerSlog) Error(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func (l *LoggerSlog) Fatal(msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
	os.Exit(1)
}

func (l *LoggerSlog) GetLogger() any {
	return l.logger
}
