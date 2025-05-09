package app

import "context"

type UseCase0[In any] interface {
	Handle(ctx context.Context, in In) error
	GetName() string
}

type UseCase1[In any, Out any] interface {
	Handle(ctx context.Context, in In) (Out, error)
	GetName() string
}
