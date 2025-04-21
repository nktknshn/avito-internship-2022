package logging

import (
	"context"
)

type Logger interface {
	InitLogger(ctx context.Context, args ...interface{})
	GetLogger() any
	Debug(ctx context.Context, msg string, args ...interface{})
	Info(ctx context.Context, msg string, args ...interface{})
	Warn(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, msg string, args ...interface{})
	Fatal(ctx context.Context, msg string, args ...interface{})
}
