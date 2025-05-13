package use_cases_test

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *UseCasesSuiteIntegrationTest) TestReportTransactions_Success() {
	acc1 := s.newAccountSaved(func(a *domainAccount.Account) {
		a.BalanceDeposit(fixtures.AmountPositive100)
	})

	acc2 := s.newAccountSaved(func(a *domainAccount.Account) {
		a.UserID = fixtures.UserID_2
		a.BalanceDeposit(fixtures.AmountPositive100)
	})

	s.Require().NoError(
		s.deposit.Handle(s.Context(), fixtures.InDeposit100),
	)

	s.Require().NoError(
		s.reserve.Handle(s.Context(), fixtures.InReserve100),
	)

	s.Require().NoError(
		s.reserveConfirm.Handle(s.Context(), fixtures.InReserveConfirm100),
	)

	trIn, err := transfer.NewInFromValues(
		acc1.ID.Value(),
		acc2.ID.Value(),
		fixtures.AmountPositive100.Value(),
	)

	s.Require().NoError(err)
	s.Require().NoError(
		s.transfer.Handle(s.Context(), trIn),
	)

}
