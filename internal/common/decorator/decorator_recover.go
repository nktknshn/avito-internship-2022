//nolint:nonamedreturns // используем в defer
package decorator

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/common/errors"
)

type RecorverHandler = func(ctx context.Context, err error) (errRecovered error)

type Decorator1Recover[T any, R any] struct {
	Base           UseCase1Handler[T, R]
	RecoverHandler RecorverHandler
}

func (d *Decorator1Recover[T, R]) Handle(ctx context.Context, in T) (out R, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = d.RecoverHandler(ctx, errors.NewErrPanic(r))
		}
	}()
	return d.Base.Handle(ctx, in)
}

func (d *Decorator1Recover[T, R]) GetName() string {
	return d.Base.GetName()
}

func (d *Decorator1Recover[T, R]) GetBase() UseCase1Handler[T, R] {
	return d.Base
}

type Decorator0Recover[T any] struct {
	Base           UseCase0Handler[T]
	RecoverHandler RecorverHandler
}

func (d *Decorator0Recover[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = d.RecoverHandler(ctx, errors.NewErrPanic(r))
		}
	}()
	return d.Base.Handle(ctx, in)
}

func (d *Decorator0Recover[T]) GetName() string {
	return d.Base.GetName()
}

func (d *Decorator0Recover[T]) GetBase() UseCase0Handler[T] {
	return d.Base
}
