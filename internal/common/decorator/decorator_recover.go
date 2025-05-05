package decorator

import (
	"context"
	"fmt"
)

type RecorverHandler = func(ctx context.Context, err error) (errRecovered error)

type Decorator1Recover[T any, R any] struct {
	base           UseCase1Handler[T, R]
	recoverHandler RecorverHandler
}

type ErrPanic struct {
	Arg any
	err error
}

func NewErrPanic(arg any) *ErrPanic {
	if e, ok := arg.(error); ok {
		return &ErrPanic{Arg: arg, err: e}
	}
	return &ErrPanic{Arg: arg}
}

func (p *ErrPanic) Error() string {
	switch e := p.Arg.(type) {
	case error:
		return fmt.Sprintf("panic: %v", e.Error())
	default:
		return fmt.Sprintf("panic: %v", e)
	}
}

func (p *ErrPanic) Unwrap() error {
	return p.err
}

func (d *Decorator1Recover[T, R]) Handle(ctx context.Context, in T) (out R, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = d.recoverHandler(ctx, NewErrPanic(r))
		}
	}()
	return d.base.Handle(ctx, in)
}

func (d *Decorator1Recover[T, R]) GetName() string {
	return d.base.GetName()
}

func (d *Decorator1Recover[T, R]) GetBase() UseCase1Handler[T, R] {
	return d.base
}

type Decorator0Recover[T any] struct {
	base           UseCase0Handler[T]
	recoverHandler RecorverHandler
}

func (d *Decorator0Recover[T]) Handle(ctx context.Context, in T) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = d.recoverHandler(ctx, NewErrPanic(r))
		}
	}()
	return d.base.Handle(ctx, in)
}

func (d *Decorator0Recover[T]) GetName() string {
	return d.base.GetName()
}

func (d *Decorator0Recover[T]) GetBase() UseCase0Handler[T] {
	return d.base
}
