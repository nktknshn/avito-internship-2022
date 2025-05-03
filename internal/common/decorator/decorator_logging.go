package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
)

type DecoratorCommandLogging[T any] struct {
	base   UseCaseCommandHandler[T]
	logger logging.Logger
}

func (d *DecoratorCommandLogging[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if err != nil {
			d.logger.Error(ctx, d.base.GetName(), "use_case", d.base.GetName(), "error", err)
		}
	}()
	d.logger.Info(ctx, d.base.GetName(), "use_case", d.base.GetName(), "in", in)
	return d.base.Handle(ctx, in)
}

func (d *DecoratorCommandLogging[T]) GetName() string {
	return d.base.GetName()
}

type DecoratorQueryLogging[T any, R any] struct {
	base   UseCaseQueryHandler[T, R]
	logger logging.Logger
}

func (d *DecoratorQueryLogging[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	defer func() {
		if err != nil {
			d.logger.Error(ctx, d.base.GetName(), "use_case", d.base.GetName(), "error", err)
		}
	}()
	d.logger.Info(ctx, d.base.GetName(), "use_case", d.base.GetName(), "in", in)
	return d.base.Handle(ctx, in)
}

func (d *DecoratorQueryLogging[T, R]) GetName() string {
	return d.base.GetName()
}
