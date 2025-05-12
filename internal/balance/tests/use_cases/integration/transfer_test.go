package use_cases_test

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func (s *UseCasesSuiteIntegrationTest) TestTransfer_Success() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID_2
	})

	in := must.Must(transfer.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.UserID_2_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.transfer.Handle(context.Background(), in)
	s.Require().NoError(err)

	fromAcc, err := s.accountsRepo.GetByUserID(context.Background(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), fromAcc.Balance.GetAvailable().Value())
	s.Require().Equal(int64(0), fromAcc.Balance.GetReserved().Value())

	toAcc, err := s.accountsRepo.GetByUserID(context.Background(), fixtures.UserID_2)
	s.Require().NoError(err)
	s.Require().Equal(fixtures.AmountPositive100_i64, toAcc.Balance.GetAvailable().Value())
	s.Require().Equal(int64(0), toAcc.Balance.GetReserved().Value())
}

func (s *UseCasesSuiteIntegrationTest) TestTransfer_AccountNotFound_To() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID
	})

	in := must.Must(transfer.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.UserID_2_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.transfer.Handle(context.Background(), in)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)

}

func (s *UseCasesSuiteIntegrationTest) TestTransfer_AccountNotFound_From() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID_2
	})

	in := must.Must(transfer.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.UserID_2_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.transfer.Handle(context.Background(), in)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestTransfer_SameAccount() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID
	})

	in := must.Must(transfer.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.UserID_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.transfer.Handle(context.Background(), in)
	s.Require().ErrorIs(err, domainAccount.ErrSameAccount)
}

func (s *UseCasesSuiteIntegrationTest) TestTransfer_InsufficientBalance() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive50))
	})

	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID_2
	})

	in := must.Must(transfer.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.UserID_2_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.transfer.Handle(context.Background(), in)
	s.Require().ErrorIs(err, domainAccount.ErrInsufficientAvailableBalance)
}

func (s *UseCasesSuiteIntegrationTest) TestTransfer_DoubleTransaction() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID_2
	})

	in := must.Must(transfer.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.UserID_2_i64,
		fixtures.AmountPositive100_i64,
	))

	workers := 20

	wg := sync.WaitGroup{}
	lock := make(chan struct{})
	errorCount := atomic.Int32{}
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-lock
			err := s.transfer.Handle(context.Background(), in)
			if err != nil {
				errorCount.Add(1)
				s.Require().ErrorIs(err, domainAccount.ErrInsufficientAvailableBalance)
			}
		}()
	}

	close(lock)
	wg.Wait()

	s.Require().Equal(int32(workers-1), errorCount.Load())
}
