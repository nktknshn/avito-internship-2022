package transfer

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

type TransferUseCase struct {
	trm             trm.Manager
	accountRepo     domainAccount.AccountRepository
	transactionRepo domainTransaction.TransactionRepository
}

func New(
	trm trm.Manager,
	accountRepo domainAccount.AccountRepository,
	transactionsRepo domainTransaction.TransactionRepository,
) *TransferUseCase {

	if trm == nil {
		panic("trm is nil")
	}

	if accountRepo == nil {
		panic("accountRepo is nil")
	}
	if transactionsRepo == nil {
		panic("transactionsRepo is nil")
	}

	return &TransferUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

func (u *TransferUseCase) Handle(ctx context.Context, in In) error {

	err := u.trm.Do(ctx, func(ctx context.Context) error {

		if in.FromUserID == in.ToUserID {
			return domainAccount.ErrSameAccount
		}

		fromAcc, err := u.accountRepo.GetByUserID(ctx, in.FromUserID)
		if err != nil {
			return err
		}

		toAcc, err := u.accountRepo.GetByUserID(ctx, in.ToUserID)
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

func (u *TransferUseCase) GetName() string {
	return use_cases.NameTransfer
}
