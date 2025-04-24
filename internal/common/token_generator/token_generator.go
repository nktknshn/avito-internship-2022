package token_generator

import (
	"context"
)

type TokenGenerator[T any] interface {
	GenerateToken(ctx context.Context, claims T) (string, error)
}
