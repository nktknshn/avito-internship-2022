//nolint:nonamedreturns // используем в defer
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

func (h recoverHandler) Handle(_ context.Context, err error) (errRecovered error) {
	h.logger.Error("panic recovered", "error", err)
	return err
}

func Decorate0[T any](
	base UseCase0Handler[T],
	metrics metrics.Metrics,
	logger logging.Logger,
) UseCase0Handler[T] {
	return &Decorator0Logging[T]{
		Base: &Decorator0Metrics[T]{
			Base: &Decorator0Recover[T]{
				Base:           base,
				RecoverHandler: recoverHandler{logger}.Handle,
			},
			Metrics: metrics,
		},
		Logger: logger,
	}
}

func Decorate1[T any, R any](
	base UseCase1Handler[T, R],
	metrics metrics.Metrics,
	logger logging.Logger,
) UseCase1Handler[T, R] {
	return &Decorator1Logging[T, R]{
		Base: &Decorator1Metrics[T, R]{
			Base: &Decorator1Recover[T, R]{
				Base:           base,
				RecoverHandler: recoverHandler{logger}.Handle,
			},
			Metrics: metrics,
		},
		Logger: logger,
	}
}
