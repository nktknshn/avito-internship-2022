package account_create

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

type AccountCreateUseCase struct {
	userAccountRepo domain.AccountRepository
}

func NewAccountCreateUseCase(userAccountRepo domain.AccountRepository) *AccountCreateUseCase {

	if userAccountRepo == nil {
		panic("userAccountRepo is nil")
	}

	return &AccountCreateUseCase{
		userAccountRepo: userAccountRepo,
	}
}

type In struct {
	UserID domain.UserID
}

func (u *AccountCreateUseCase) Handle(ctx context.Context, in In) error {
	return nil
}
