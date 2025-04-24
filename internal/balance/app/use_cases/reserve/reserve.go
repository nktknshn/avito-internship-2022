package reserve

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type ReserveUseCase struct {
	trm             trm.Manager
	accountRepo     domainAccount.AccountRepository
	transactionRepo domainTransaction.TransactionRepository
}

func New(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionRepo domainTransaction.TransactionRepository,
) *ReserveUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	if transactionRepo == nil {
		panic("transactionRepo == nil")
	}

	return &ReserveUseCase{
		trm,
		accountRepo,
		transactionRepo,
	}
}

func (u *ReserveUseCase) Handle(ctx context.Context, in In) error {

	// а если canceled, то OrderID новый?
	// если есть резерв с OrderID и статус не canceled, то ошибка

	err := u.trm.Do(ctx, func(ctx context.Context) error {

		acc, err := u.accountRepo.GetByUserID(ctx, in.userID)

		if err != nil {
			return err
		}

		orderTransactions, err := u.transactionRepo.GetTransactionSpendByOrderID(ctx, in.userID, in.orderID)

		if err != nil {
			return err
		}

		for _, transaction := range orderTransactions {
			// если существует транзакция с таким OrderID и статус не canceled, то ошибка
			if transaction.Status != domainTransaction.TransactionSpendStatusCanceled {
				return domainTransaction.ErrTransactionAlreadyExists
			}
		}

		err = acc.BalanceReserve(in.amount)

		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, acc)

		if err != nil {
			return err
		}

		transaction, err := domainTransaction.NewTransactionSpendReserved(
			acc.ID,
			in.userID,
			in.orderID,
			in.productID,
			in.amount,
			time.Now(),
		)

		if err != nil {
			return err
		}

		_, err = u.transactionRepo.SaveTransactionSpend(ctx, transaction)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
