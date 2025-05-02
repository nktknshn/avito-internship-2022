package get_balance

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type GetBalanceUseCase struct {
	trm         trm.Manager
	accountRepo domainAccount.AccountRepository
}

func New(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
) *GetBalanceUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	return &GetBalanceUseCase{
		trm,
		accountRepo,
	}
}

func (u *GetBalanceUseCase) Handle(ctx context.Context, in In) (Out, error) {

	acc, err := u.accountRepo.GetByUserID(ctx, in.userID)

	if err != nil {
		return Out{}, err
	}

	return Out{
		Available: acc.Balance.GetAvailable(),
		Reserved:  acc.Balance.GetReserved(),
	}, nil
}

func (u *GetBalanceUseCase) GetName() string {
	return use_cases.NameGetBalance
}
