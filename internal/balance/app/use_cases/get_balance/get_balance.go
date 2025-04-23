package get_balance

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
)

type getBalanceUseCase struct {
	trm         trm.Manager
	accountRepo domainAccount.AccountRepository
}

func New(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
) *getBalanceUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	return &getBalanceUseCase{
		trm,
		accountRepo,
	}
}

type Out struct {
	Available int64
	Reserved  int64
}

func (u *getBalanceUseCase) Handle(ctx context.Context, in In) (Out, error) {

	acc, err := u.accountRepo.GetByUserID(ctx, in.UserID)

	if err != nil {
		return Out{}, err
	}

	return Out{
		Available: acc.Balance.GetAvailable().Value(),
		Reserved:  acc.Balance.GetReserved().Value(),
	}, nil
}
