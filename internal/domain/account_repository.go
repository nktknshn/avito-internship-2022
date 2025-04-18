package domain

import (
	"context"
	"errors"
)

var (
	ErrAccountNotFound = errors.New("account not found")
)

type AccountRepository interface {
	Save(ctx context.Context, account *Account) (*Account, error)
	GetByUserID(ctx context.Context, userID UserID) (*Account, error)
}
