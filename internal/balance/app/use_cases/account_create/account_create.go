package account_create

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type AccountCreateUseCase struct {
	userAccountRepo domainAccount.AccountRepository
}

func NewAccountCreateUseCase(userAccountRepo domainAccount.AccountRepository) *AccountCreateUseCase {

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
	acc, err := domainAccount.NewAccount(in.UserID)
	if err != nil {
		return err
	}

	_, err = u.userAccountRepo.Save(ctx, acc)
	return err
}

// func (u *AccountCreateUseCase) GetName() string {
// 	return use_cases.AccountCreate
// }
