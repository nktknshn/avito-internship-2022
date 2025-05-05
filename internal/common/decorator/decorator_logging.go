package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
)

type Decorator0Logging[T any] struct {
	base   UseCase0Handler[T]
	logger logging.Logger
}

func (d *Decorator0Logging[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if err != nil {
			d.logger.Error(ctx, d.base.GetName(), "use_case", d.base.GetName(), "error", err)
		}
	}()
	d.logger.Info(ctx, d.base.GetName(), "use_case", d.base.GetName(), "in", in)
	return d.base.Handle(ctx, in)
}

func (d *Decorator0Logging[T]) GetName() string {
	return d.base.GetName()
}

type Decorator1Logging[T any, R any] struct {
	base   UseCase1Handler[T, R]
	logger logging.Logger
}

func (d *Decorator1Logging[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	defer func() {
		if err != nil {
			d.logger.Error(ctx, d.base.GetName(), "use_case", d.base.GetName(), "error", err)
		}
	}()
	d.logger.Info(ctx, d.base.GetName(), "use_case", d.base.GetName(), "in", in)
	return d.base.Handle(ctx, in)
}

func (d *Decorator1Logging[T, R]) GetName() string {
	return d.base.GetName()
}
