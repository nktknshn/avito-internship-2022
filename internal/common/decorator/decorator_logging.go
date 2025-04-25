package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
)

type DecoratorCommandLogging[T any] struct {
	base       UseCaseCommandHandler[T]
	logger     logging.Logger
	methodName string
}

func (d *DecoratorCommandLogging[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if err != nil {
			d.logger.Error(ctx, d.methodName, "error", err)
		}
	}()
	d.logger.Info(ctx, d.methodName, "in", in)
	return d.base.Handle(ctx, in)
}

type DecoratorQueryLogging[T any, R any] struct {
	base       UseCaseQueryHandler[T, R]
	logger     logging.Logger
	methodName string
}

func (d *DecoratorQueryLogging[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	defer func() {
		if err != nil {
			d.logger.Error(ctx, d.methodName, "error", err)
		}
	}()
	d.logger.Info(ctx, d.methodName, "in", in)
	return d.base.Handle(ctx, in)
}
