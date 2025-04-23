package account

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrAccountNotFound = domainError.New("account not found")
)

type AccountRepository interface {
	Save(ctx context.Context, account *Account) (*Account, error)
	GetByUserID(ctx context.Context, userID domain.UserID) (*Account, error)
	GetByAccountID(ctx context.Context, accountID AccountID) (*Account, error)
}
