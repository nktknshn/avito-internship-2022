package domain

import (
	"context"
)

type AccountRepository interface {
	Save(ctx context.Context, account *Account) (*Account, error)
	GetByUserID(ctx context.Context, userID UserID) (*Account, error)
}
