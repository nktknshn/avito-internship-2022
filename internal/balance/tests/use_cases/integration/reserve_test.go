package use_cases_test

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *UseCasesSuiteIntegrationTest) TestReserve_Success() {
	acc := s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.GetByUserID(context.Background(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(fixtures.AmountPositive100_i64, acc.Balance.GetReserved().Value())
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())

	// транзакция должна быть создана
	transactions, err := s.transactionsRepo.GetTransactionSpendByOrderID(context.Background(), fixtures.UserID, fixtures.OrderID)
	s.Require().NoError(err)
	s.Require().Equal(1, len(transactions))

	transaction := transactions[0]
	s.Require().Equal(domainTransaction.TransactionSpendStatusReserved, transaction.Status)
	s.Require().Equal(fixtures.AmountPositive100_i64, transaction.Amount.Value())

}

func (s *UseCasesSuiteIntegrationTest) TestReserve_AccountNotFound() {
	err := s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_InsufficientBalance() {
	acc := s.newAccount()
	acc, err := s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	err = s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainAccount.ErrInsufficientAvailableBalance)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_AlreadyPaid() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().NoError(err)

	err = s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm100)
	s.Require().NoError(err)

	err = s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyPaid)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_AlreadyExists() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().NoError(err)

	err = s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyExists)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_DoubleReserve() {
	acc := s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	workers := 20

	wg := sync.WaitGroup{}
	lock := make(chan struct{})
	errorCount := atomic.Int32{}
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-lock
			err := s.reserve.Handle(context.Background(), fixtures.InReserve100)
			if err != nil {
				errorCount.Add(1)
				isPaidOrExists := errors.Is(err,
					domainTransaction.ErrTransactionAlreadyPaid) ||
					errors.Is(err, domainTransaction.ErrTransactionAlreadyExists)
				s.Require().True(isPaidOrExists)
			}
		}()
	}

	close(lock)
	wg.Wait()

	s.Require().Equal(int32(workers-1), errorCount.Load())

	acc, err := s.accountsRepo.GetByUserID(context.Background(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())
}
