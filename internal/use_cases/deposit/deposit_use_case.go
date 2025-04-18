package deposit

import (
	"context"
	"errors"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

type DepositUseCase struct {
	trm              trm.Manager
	accountRepo      domain.AccountRepository
	transactionsRepo domain.TransactionRepository
}

func NewDepositUseCase(
	trm trm.Manager,
	accountRepo domain.AccountRepository,
	transactionsRepo domain.TransactionRepository,
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

type In struct {
	UserID domain.UserID
	Amount domain.AmountPositive
	Source domain.DepositSource
}

func NewInFromValues(userID int64, amount int64, source string) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}

	_source, err := domain.NewDepositSource(source)
	if err != nil {
		return In{}, err
	}

	_amount, err := domain.NewAmountPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{
		UserID: _userID,
		Source: _source,
		Amount: _amount,
	}, nil
}

// getAccountCreating returns an account for the user.
// If the account does not exist, it creates a new one.
func (u *DepositUseCase) getAccountCreating(ctx context.Context, userID domain.UserID) (*domain.Account, error) {

	acc, err := u.accountRepo.GetByUserID(ctx, userID)

	if errors.Is(err, domain.ErrAccountNotFound) {
		newAccount, err := domain.NewAccount(userID)

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

		transaction, err := domain.NewTransactionDeposit(
			acc.ID,
			in.UserID,
			in.Source,
			in.Amount,
			time.Now(),
		)

		if err != nil {
			return err
		}

		_, err = u.transactionsRepo.SaveTransactionDeposit(ctx, transaction)

		if err != nil {
			return err
		}

		err = acc.BalanceDeposit(in.Amount)

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
