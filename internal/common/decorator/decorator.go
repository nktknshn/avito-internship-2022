package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
)

type UseCase0Handler[T any] interface {
	Handle(ctx context.Context, in T) error
	GetName() string
}

type UseCase1Handler[T any, R any] interface {
	Handle(ctx context.Context, in T) (R, error)
	GetName() string
}

type recoverHandler struct {
	logger logging.Logger
}

func (h recoverHandler) Handle(ctx context.Context, err error) (errRecovered error) {
	h.logger.Error("panic recovered", "error", err)
	return err
}

func Decorate0[T any](
	base UseCase0Handler[T],
	metrics metrics.Metrics,
	logger logging.Logger,
) UseCase0Handler[T] {
	return &Decorator0Logging[T]{
		base: &Decorator0Metrics[T]{
			base: &Decorator0Recover[T]{
				base:           base,
				recoverHandler: recoverHandler{logger}.Handle,
			},
			metrics: metrics,
		},
		logger: logger,
	}
}

func Decorate1[T any, R any](
	base UseCase1Handler[T, R],
	metrics metrics.Metrics,
	logger logging.Logger,
) UseCase1Handler[T, R] {
	return &Decorator1Logging[T, R]{
		base: &Decorator1Metrics[T, R]{
			base: &Decorator1Recover[T, R]{
				base:           base,
				recoverHandler: recoverHandler{logger}.Handle,
			},
			metrics: metrics,
		},
		logger: logger,
	}
}
