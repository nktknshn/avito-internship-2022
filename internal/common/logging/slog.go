package logging

import (
	"context"
	"log/slog"
	"os"
)

type LoggerSlog struct {
	logger *slog.Logger
}

func NewSlog() *LoggerSlog {
	return &LoggerSlog{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

// InitLogger
func (l *LoggerSlog) InitLogger(_ context.Context, args ...interface{}) {
	l.logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func (l *LoggerSlog) Debug(_ context.Context, msg string, args ...interface{}) {
	l.logger.Debug(msg, args...)
}

func (l *LoggerSlog) Info(_ context.Context, msg string, args ...interface{}) {
	l.logger.Info(msg, args...)
}

func (l *LoggerSlog) Warn(_ context.Context, msg string, args ...interface{}) {
	l.logger.Warn(msg, args...)
}

func (l *LoggerSlog) Error(_ context.Context, msg string, args ...interface{}) {
	l.logger.Error(msg, args...)
}

func (l *LoggerSlog) Fatal(_ context.Context, msg string, args ...interface{}) {

	l.logger.Error(msg, args...)
	os.Exit(1)
}

func (l *LoggerSlog) GetLogger() any {
	return l.logger
}
