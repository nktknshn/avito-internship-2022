package use_cases_test

import (
	"context"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_Success() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	s.Require().NoError(s.reserve.Handle(context.Background(), fixtures.InReserve100))

	err := s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm100)
	s.Require().NoError(err)

	acc, err := s.accountsRepo.GetByUserID(context.Background(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), acc.Balance.GetReserved().Value())
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())

	transactions, err := s.transactionsRepo.GetTransactionSpendByOrderID(context.Background(), fixtures.UserID, fixtures.OrderID)
	s.Require().NoError(err)
	s.Require().Equal(1, len(transactions))

	transaction := transactions[0]
	s.Require().Equal(domainTransaction.TransactionSpendStatusConfirmed, transaction.Status)
	s.Require().Equal(fixtures.AmountPositive100_i64, transaction.Amount.Value())
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_AccountNotFound() {
	err := s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm100)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionNotFound() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionAmountMismatch() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	s.Require().NoError(s.reserve.Handle(context.Background(), fixtures.InReserve100))

	err := s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm50)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAmountMismatch)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionProductIDMismatch() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	s.Require().NoError(s.reserve.Handle(context.Background(), fixtures.InReserve100))

	copyIn := fixtures.InReserveConfirm100
	copyIn.ProductID = 6666

	err := s.reserveConfirm.Handle(context.Background(), copyIn)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionProductIDMismatch)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionAlreadyPaid() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(context.Background(), fixtures.InReserve100)
	s.Require().NoError(err)

	err = s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm100)
	s.Require().NoError(err)

	err = s.reserveConfirm.Handle(context.Background(), fixtures.InReserveConfirm100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyPaid)
}
