package use_cases_test

import (
	"errors"
	"sync"
	"sync/atomic"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *UseCasesSuiteIntegrationTest) TestReserve_Success() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().NoError(err)

	acc, err := s.accountsRepo.GetByUserID(s.Context(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(fixtures.AmountPositive100_i64, acc.Balance.GetReserved().Value())
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())

	// транзакция должна быть создана
	transactions, err := s.transactionsRepo.GetTransactionSpendByOrderID(s.Context(), fixtures.UserID, fixtures.OrderID)
	s.Require().NoError(err)
	s.Require().Equal(1, len(transactions))

	transaction := transactions[0]
	s.Require().Equal(domainTransaction.TransactionSpendStatusReserved, transaction.Status)
	s.Require().Equal(fixtures.AmountPositive100_i64, transaction.Amount.Value())

}

func (s *UseCasesSuiteIntegrationTest) TestReserve_AccountNotFound() {
	err := s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_InsufficientBalance() {
	acc := s.newAccount()
	_, err := s.accountsRepo.Save(s.Context(), acc)
	s.Require().NoError(err)

	err = s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainAccount.ErrInsufficientAvailableBalance)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_AlreadyPaid() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().NoError(err)

	err = s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100)
	s.Require().NoError(err)

	err = s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyPaid)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_AlreadyExists() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().NoError(err)

	err = s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyExists)
}

func (s *UseCasesSuiteIntegrationTest) TestReserve_DoubleReserve() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	workers := 20

	wg := sync.WaitGroup{}
	lock := make(chan struct{})
	errorCount := atomic.Int32{}
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-lock
			err := s.reserve.Handle(s.Context(), fixtures.InReserve100)
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

	acc, err := s.accountsRepo.GetByUserID(s.Context(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())
}
