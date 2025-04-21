package reserve_confirm

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type reserveConfirmUseCase struct {
	trm              trm.Manager
	accountRepo      domainAccount.AccountRepository
	transactionsRepo domainTransaction.TransactionRepository
}

func NewReserveConfirmUseCase(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionsRepo domainTransaction.TransactionRepository,
) *reserveConfirmUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	if transactionsRepo == nil {
		panic("transactionsRepo == nil")
	}

	return &reserveConfirmUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

func (u *reserveConfirmUseCase) Handle(ctx context.Context, in In) error {
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
			return err
		}

		err = transaction.Confirm(time.Now())

		if err != nil {
			return err
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
