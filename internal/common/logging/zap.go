package logging

import (
	"go.uber.org/zap"
)

type LoggerZap struct {
	logger *zap.Logger
}

func NewLoggerZap() *LoggerZap {
	return &LoggerZap{logger: zap.NewExample()}
}
