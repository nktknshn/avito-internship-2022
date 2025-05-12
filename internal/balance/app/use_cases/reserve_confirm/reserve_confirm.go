package reserve_confirm

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/pkg/errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type ReserveConfirmUseCase struct {
	trm              trm.Manager
	accountRepo      domainAccount.AccountRepository
	transactionsRepo domainTransaction.TransactionRepository
}

func New(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionsRepo domainTransaction.TransactionRepository,
) *ReserveConfirmUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	if transactionsRepo == nil {
		panic("transactionsRepo == nil")
	}

	return &ReserveConfirmUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

func (u *ReserveConfirmUseCase) Handle(ctx context.Context, in In) error {
	// если amount не равен сумме резерва, то ошибка
	err := u.trm.Do(ctx, func(ctx context.Context) error {
		acc, err := u.accountRepo.GetByUserID(ctx, in.UserID)

		if err != nil {
			return err
		}

		orderTransactions, err := u.transactionsRepo.GetTransactionSpendByOrderID(ctx, in.UserID, in.OrderID)

		if err != nil {
			return err
		}

		var transaction *domainTransaction.TransactionSpend

		for _, transaction = range orderTransactions {
			if transaction.Status == domainTransaction.TransactionSpendStatusConfirmed {
				return domainTransaction.ErrTransactionAlreadyPaid
			}

			if transaction.ProductID != in.ProductID {
				return domainTransaction.ErrTransactionProductIDMismatch
			}

			if transaction.Status == domainTransaction.TransactionSpendStatusReserved {
				break
			}
		}

		if transaction == nil {
			return domainTransaction.ErrTransactionNotFound
		}

		// возможно имплементировать как стратегию на уровне моделей домена
		if transaction.Amount != in.Amount {
			return domainTransaction.ErrTransactionAmountMismatch
		}

		err = acc.BalanceReserveConfirm(in.Amount)

		if err != nil {
			return errors.Wrap(err, "ReserveConfirmUseCase.BalanceReserveConfirm")
		}

		err = transaction.Confirm(time.Now())

		if err != nil {
			return errors.Wrap(err, "ReserveConfirmUseCase.transaction.Confirm")
		}

		_, err = u.transactionsRepo.SaveTransactionSpend(ctx, transaction)

		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, acc)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (u *ReserveConfirmUseCase) GetName() string {
	return use_cases.NameReserveConfirm
}
