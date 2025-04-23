package deposit

import (
	"context"
	"errors"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"

	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type DepositUseCase struct {
	trm              trm.Manager
	accountRepo      domainAccount.AccountRepository
	transactionsRepo domainTransaction.TransactionRepository
}

func New(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionsRepo domainTransaction.TransactionRepository,
) *DepositUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("userRepo == nil")
	}

	if transactionsRepo == nil {
		panic("transactionsRepo == nil")
	}

	return &DepositUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

// getAccountCreating returns an account for the user.
// If the account does not exist, it creates a new one.
func (u *DepositUseCase) getAccountCreating(ctx context.Context, userID domain.UserID) (*domainAccount.Account, error) {

	acc, err := u.accountRepo.GetByUserID(ctx, userID)

	if errors.Is(err, domainAccount.ErrAccountNotFound) {
		newAccount, err := domainAccount.NewAccount(userID)

		if err != nil {
			return nil, err
		}

		acc, err = u.accountRepo.Save(ctx, newAccount)

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (u *DepositUseCase) Handle(ctx context.Context, in In) error {

	err := u.trm.Do(ctx, func(ctx context.Context) error {

		acc, err := u.getAccountCreating(ctx, in.UserID)

		if err != nil {
			return nil
		}

		transaction, err := domainTransaction.NewTransactionDeposit(
			acc.ID,
			in.UserID,
			in.Source,
			in.Amount,
			time.Now(),
		)

		if err != nil {
			return err
		}

		err = acc.BalanceDeposit(in.Amount)

		if err != nil {
			return err
		}

		_, err = u.transactionsRepo.SaveTransactionDeposit(ctx, transaction)

		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, acc)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
