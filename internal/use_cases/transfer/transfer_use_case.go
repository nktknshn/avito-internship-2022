package transfer

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/domain/transaction"
)

type transferUseCase struct {
	trm             trm.Manager
	accountRepo     domainAccount.AccountRepository
	transactionRepo domainTransaction.TransactionRepository
}

func NewTransferUseCase(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionsRepo domainTransaction.TransactionRepository,
) *transferUseCase {

	if trm == nil {
		panic("trm is nil")
	}

	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	if transactionsRepo == nil {
		panic("transactionsRepo is nil")
	}

	return &transferUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

func (u *transferUseCase) Handle(ctx context.Context, in In) error {

	err := u.trm.Do(ctx, func(ctx context.Context) error {

		if in.From == in.To {
			return domainAccount.ErrSameAccount
		}

		fromAcc, err := u.accountRepo.GetByAccountID(ctx, in.From)
		if err != nil {
			return err
		}

		toAcc, err := u.accountRepo.GetByAccountID(ctx, in.To)
		if err != nil {
			return err
		}

		err = fromAcc.Transfer(toAcc, in.Amount)

		if err != nil {
			return err
		}

		transaction, err := domainTransaction.NewTransactionTransfer(
			fromAcc.ID,
			toAcc.ID,
			in.Amount,
			time.Now(),
		)
		if err != nil {
			return err
		}

		_, err = u.transactionRepo.SaveTransactionTransfer(ctx, transaction)
		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, fromAcc)
		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, toAcc)
		if err != nil {
			return err
		}

		return nil

	})

	return err
}
