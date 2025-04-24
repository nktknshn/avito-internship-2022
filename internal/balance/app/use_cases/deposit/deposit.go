package deposit

import (
	"context"
	"time"

	"github.com/pkg/errors"

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
			return nil, errors.Wrap(err, "DepositUseCase.getAccountCreating.NewAccount")
		}

		acc, err = u.accountRepo.Save(ctx, newAccount)

		if err != nil {
			return nil, errors.Wrap(err, "DepositUseCase.getAccountCreating.Save")
		}

		return acc, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, "DepositUseCase.getAccountCreating")
	}

	return acc, nil
}

func (u *DepositUseCase) Handle(ctx context.Context, in In) error {

	err := u.trm.Do(ctx, func(ctx context.Context) error {

		acc, err := u.getAccountCreating(ctx, in.userID)

		if err != nil {
			return errors.Wrap(err, "DepositUseCase.Handle.getAccountCreating")
		}

		transaction, err := domainTransaction.NewTransactionDeposit(
			acc.ID,
			in.userID,
			in.source,
			in.amount,
			time.Now(),
		)

		if err != nil {
			return errors.Wrap(err, "DepositUseCase.Handle.NewTransactionDeposit")
		}

		err = acc.BalanceDeposit(in.amount)

		if err != nil {
			return errors.Wrap(err, "DepositUseCase.Handle.BalanceDeposit")
		}

		_, err = u.transactionsRepo.SaveTransactionDeposit(ctx, transaction)

		if err != nil {
			return errors.Wrap(err, "DepositUseCase.Handle.SaveTransactionDeposit")
		}

		_, err = u.accountRepo.Save(ctx, acc)

		if err != nil {
			return errors.Wrap(err, "DepositUseCase.Handle.SaveAccount")
		}

		return nil
	})

	if err != nil {
		return errors.Wrap(err, "DepositUseCase.Handle")
	}

	return nil
}
