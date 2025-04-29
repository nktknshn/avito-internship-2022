package use_cases_test

import (
	"context"
	"sync"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func (s *SuiteTest) TestReserve_Success() {
	acc := s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	in := must.Must(reserve.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.ProductID_i64,
		fixtures.ProductTitle_str,
		fixtures.OrderID_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.reserve.Handle(context.Background(), in)
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

func (s *SuiteTest) TestReserve_AccountNotFound() {

	in := must.Must(reserve.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.ProductID_i64,
		fixtures.ProductTitle_str,
		fixtures.OrderID_i64,
		fixtures.AmountPositive100_i64,
	))

	err := s.reserve.Handle(context.Background(), in)

	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *SuiteTest) TestReserve_InsufficientBalance() {
	acc := s.newAccount()
	acc, err := s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	in := must.Must(reserve.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.ProductID_i64,
		fixtures.ProductTitle_str,
		fixtures.OrderID_i64,
		fixtures.AmountPositive100_i64,
	))

	err = s.reserve.Handle(context.Background(), in)
	s.Require().ErrorIs(err, domainAccount.ErrInsufficientAvailableBalance)
}

func (s *SuiteTest) TestReserve_DoubleReserve() {
	acc := s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	in := must.Must(reserve.NewInFromValues(
		fixtures.UserID_i64,
		fixtures.ProductID_i64,
		fixtures.ProductTitle_str,
		fixtures.OrderID_i64,
		fixtures.AmountPositive100_i64,
	))

	wg := sync.WaitGroup{}
	lock := make(chan struct{})
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-lock
			err := s.reserve.Handle(context.Background(), in)
			if err != nil {
				s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyExists)
			}
		}()
	}

	close(lock)
	wg.Wait()

	acc, err := s.accountsRepo.GetByUserID(context.Background(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())
}
