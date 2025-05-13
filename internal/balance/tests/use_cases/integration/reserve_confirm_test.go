package use_cases_test

import (
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_Success() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	s.Require().NoError(s.reserve.Handle(s.Context(), fixtures.InReserve100))

	err := s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100)
	s.Require().NoError(err)

	acc, err := s.accountsRepo.GetByUserID(s.Context(), fixtures.UserID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), acc.Balance.GetReserved().Value())
	s.Require().Equal(int64(0), acc.Balance.GetAvailable().Value())

	transactions, err := s.transactionsRepo.GetTransactionSpendByOrderID(s.Context(), fixtures.UserID, fixtures.OrderID)
	s.Require().NoError(err)
	s.Require().Equal(1, len(transactions))

	transaction := transactions[0]
	s.Require().Equal(domainTransaction.TransactionSpendStatusConfirmed, transaction.Status)
	s.Require().Equal(fixtures.AmountPositive100_i64, transaction.Amount.Value())
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_AccountNotFound() {
	err := s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionNotFound() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionNotFound)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionAmountMismatch() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	s.Require().NoError(s.reserve.Handle(s.Context(), fixtures.InReserve100))

	err := s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm50)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAmountMismatch)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionProductIDMismatch() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	s.Require().NoError(s.reserve.Handle(s.Context(), fixtures.InReserve100))

	copyIn := fixtures.InReserveConfirm100
	copyIn.ProductID = 6666

	err := s.reserveConfirm.Handle(s.Context(), copyIn)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionProductIDMismatch)
}

func (s *UseCasesSuiteIntegrationTest) TestReserveConfirm_TransactionAlreadyPaid() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(fixtures.AmountPositive100))
	})

	err := s.reserve.Handle(s.Context(), fixtures.InReserve100)
	s.Require().NoError(err)

	err = s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100)
	s.Require().NoError(err)

	err = s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100)
	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAlreadyPaid)
}
