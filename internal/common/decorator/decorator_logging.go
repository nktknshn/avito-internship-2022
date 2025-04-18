package decorator

import "context"

type DecoratorCommandLogging[T any] struct {
	base UseCaseCommandHandler[T]
}

func (d *DecoratorCommandLogging[T]) Handle(ctx context.Context, in T) error {
	return d.base.Handle(ctx, in)
}

type DecoratorQueryLogging[T any, R any] struct {
	base UseCaseQueryHandler[T, R]
}

func (d *DecoratorQueryLogging[T, R]) Handle(ctx context.Context, in T) (R, error) {
	return d.base.Handle(ctx, in)
}
