package get_balance

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

type GetBalanceUseCase struct {
	trm         trm.Manager
	accountRepo domain.AccountRepository
}

func NewGetBalanceUseCase(
	trm trm.Manager,
	accountRepo domain.AccountRepository,
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

type In struct {
	UserID domain.UserID
}

type Out struct {
	Available int64
	Reserved  int64
}

func (u *GetBalanceUseCase) Handle(ctx context.Context, in In) (Out, error) {

	acc, err := u.accountRepo.GetByUserID(ctx, in.UserID)

	if err != nil {
		return Out{}, err
	}

	return Out{
		Available: acc.Balance.Available.Value(),
		Reserved:  acc.Balance.Reserved.Value(),
	}, nil
}
