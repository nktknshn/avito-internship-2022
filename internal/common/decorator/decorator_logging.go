//nolint:nonamedreturns // используем в defer
package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
)

type Decorator0Logging[T any] struct {
	Base   UseCase0Handler[T]
	Logger logging.Logger
}

func (d *Decorator0Logging[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if err != nil {
			d.Logger.Error(d.Base.GetName(), "use_case", d.Base.GetName(), "error", err)
		}
	}()
	d.Logger.Info(d.Base.GetName(), "use_case", d.Base.GetName(), "in", in)
	return d.Base.Handle(ctx, in)
}

func (d *Decorator0Logging[T]) GetName() string {
	return d.Base.GetName()
}

type Decorator1Logging[T any, R any] struct {
	Base   UseCase1Handler[T, R]
	Logger logging.Logger
}

func (d *Decorator1Logging[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	defer func() {
		if err != nil {
			d.Logger.Error(d.Base.GetName(), "use_case", d.Base.GetName(), "error", err)
		}
	}()
	d.Logger.Info(d.Base.GetName(), "use_case", d.Base.GetName(), "in", in)
	return d.Base.Handle(ctx, in)
}

func (d *Decorator1Logging[T, R]) GetName() string {
	return d.Base.GetName()
}
