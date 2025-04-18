package decorator

import "context"

type UseCaseCommandHandler[T any] interface {
	Handle(ctx context.Context, cmd T) error
}

type UseCaseQueryHandler[T any, R any] interface {
	Handle(ctx context.Context, cmd T) (R, error)
}
