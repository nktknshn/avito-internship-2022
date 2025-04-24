package use_cases_test

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func (s *SuiteTest) TestReserveCancel_Success() {
	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		s.Require().NoError(a.BalanceDeposit(amount100))
	})

	inReserve := must.Must(reserve.NewInFromValues(
		userID.Value(),
		orderID.Value(),
		productID.Value(),
		amount100.Value(),
	))

	s.Require().NoError(s.reserve.Handle(context.Background(), inReserve))

	inCancel := must.Must(reserve_cancel.NewInFromValues(
		userID.Value(),
		orderID.Value(),
		productID.Value(),
		amount100.Value(),
	))

	err := s.reserveCancel.Handle(context.Background(), inCancel)
	s.Require().NoError(err)

	acc, err := s.accountsRepo.GetByUserID(context.Background(), userID)
	s.Require().NoError(err)
	s.Require().Equal(int64(0), acc.Balance.GetReserved().Value())
	s.Require().Equal(amount100.Value(), acc.Balance.GetAvailable().Value())

	transactions, err := s.transactionsRepo.GetTransactionSpendByOrderID(context.Background(), userID, orderID)
	s.Require().NoError(err)
	s.Require().Equal(1, len(transactions))

	transaction := transactions[0]
	s.Require().Equal(domainTransaction.TransactionSpendStatusCanceled, transaction.Status)
	s.Require().Equal(amount100.Value(), transaction.Amount.Value())
}

func (s *SuiteTest) TestReserveCancel_AccountNotFound() {
	in := must.Must(reserve_cancel.NewInFromValues(
		userID.Value(),
		orderID.Value(),
		productID.Value(),
		amount100.Value(),
	))

	err := s.reserveCancel.Handle(context.Background(), in)

	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *SuiteTest) TestReserveCancel_TransactionNotFound() {

	_ = s.newAccountSaved()

	in := must.Must(reserve_cancel.NewInFromValues(
		userID.Value(),
		orderID.Value(),
		productID.Value(),
		amount100.Value(),
	))

	err := s.reserveCancel.Handle(context.Background(), in)

	s.Require().ErrorIs(err, domainTransaction.ErrTransactionNotFound)
}

func (s *SuiteTest) TestReserveCancel_TransactionAmountMismatch() {

	_ = s.newAccountSaved(func(a *domainAccount.Account) {
		a.BalanceDeposit(amount100)
	})

	inReserve := must.Must(reserve.NewInFromValues(
		userID.Value(),
		orderID.Value(),
		productID.Value(),
		amount100.Value(),
	))

	s.Require().NoError(s.reserve.Handle(context.Background(), inReserve))

	in := must.Must(reserve_cancel.NewInFromValues(
		userID.Value(),
		orderID.Value(),
		productID.Value(),
		amount50.Value(),
	))

	err := s.reserveCancel.Handle(context.Background(), in)

	s.Require().ErrorIs(err, domainTransaction.ErrTransactionAmountMismatch)
}
