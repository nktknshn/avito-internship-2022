package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
)

type UseCaseCommandHandler[T any] interface {
	Handle(ctx context.Context, in T) error
}

type UseCaseQueryHandler[T any, R any] interface {
	Handle(ctx context.Context, in T) (R, error)
}

func DecorateCommand[T any](
	base UseCaseCommandHandler[T],
	metrics metrics.Metrics,
	logger logging.Logger,
	methodName string,
) UseCaseCommandHandler[T] {
	return &DecoratorCommandLogging[T]{
		base: &DecoratorCommandMetrics[T]{
			base:       base,
			metrics:    metrics,
			methodName: methodName,
		},
		logger:     logger,
		methodName: methodName,
	}
}

func DecorateQuery[T any, R any](
	base UseCaseQueryHandler[T, R],
	metrics metrics.Metrics,
	logger logging.Logger,
	methodName string,
) UseCaseQueryHandler[T, R] {
	return &DecoratorQueryLogging[T, R]{
		base: &DecoratorQueryMetrics[T, R]{
			base:       base,
			metrics:    metrics,
			methodName: methodName,
		},
		logger:     logger,
		methodName: methodName,
	}
}
