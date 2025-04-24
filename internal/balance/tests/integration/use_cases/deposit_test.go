package use_cases_test

import (
	"context"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func (s *SuiteTest) TestDeposit_CreatesAccountIfNotExists() {

	in := must.Must(deposit.NewInFromValues(1, 100, ""))

	err := s.deposit.Handle(context.Background(), in)

	s.Require().NoError(err)

	rows, err := s.ExecSql("select * from accounts")
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))

	rows, err = s.ExecSql("select * from transactions_deposit")
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))
}

func (s *SuiteTest) TestDeposit_DepositsExistingAccount() {

	s.ExecSqlExpectRowsLen("select * from accounts", 0)

	acc := must.Must(domainAccount.NewAccount(1))
	acc, err := s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	s.ExecSqlExpectRowsLen("select * from accounts", 1)

	in := must.Must(deposit.NewInFromValues(1, 100, ""))

	err = s.deposit.Handle(context.Background(), in)
	s.Require().NoError(err)

	rows, err := s.ExecSql("select * from accounts")
	s.Require().NoError(err)

	s.Require().Equal(1, len(rows.Rows))
	s.Require().Equal(acc.ID.Value(), rows.Rows[0]["id"])
	s.Require().Equal(int64(100), rows.Rows[0]["balance_available"])

	rows, err = s.ExecSql("select * from transactions_deposit")
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))
}
