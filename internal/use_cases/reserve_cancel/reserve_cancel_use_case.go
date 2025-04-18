package reserve_cancel

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type ReserveCancelUseCase struct {
	trm              trm.Manager
	accountRepo      domainAccount.AccountRepository
	transactionsRepo domainTransaction.TransactionRepository
}

func NewReserveCancelUseCase(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionsRepo domainTransaction.TransactionRepository,
) *ReserveCancelUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	if transactionsRepo == nil {
		panic("transactionsRepo == nil")
	}

	return &ReserveCancelUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

func (u *ReserveCancelUseCase) Handle(ctx context.Context, in In) error {
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
			if transaction.Status == domainTransaction.TransactionSpendStatusReserved {
				break
			}
		}

		if transaction == nil {
			return domainTransaction.ErrTransactionNotFound
		}

		if transaction.Amount != in.Amount {
			return domainTransaction.ErrTransactionAmountMismatch
		}

		err = transaction.Cancel(time.Now())

		if err != nil {
			return err
		}

		_, err = u.transactionsRepo.SaveTransactionSpend(ctx, transaction)

		if err != nil {
			return err
		}

		err = acc.BalanceReserveCancel(in.Amount)

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
