package accounts_pg

import (
	"context"
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
)

func TestAccountsPg(t *testing.T) {
	s := &Suite{}
	s.SetPostgresMigrationsDir("../../../migrations/postgres")
	suite.Run(t, s)
}

type Suite struct {
	testing_pg.TestSuitePg
	accountsRepo *accounts_pg.AccountsRepository
}

func (s *Suite) SetupTest() {
	s.accountsRepo = accounts_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
}

func (s *Suite) TearDownTest() {
	s.ExecSql("delete from accounts")
}

func (s *Suite) TestSave_Create_Success() {
	acc, err := domainAccount.NewAccountFromValues(0, 1, 100, 0)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

}

func (s *Suite) TestSave_Update_Success() {
	acc, err := domainAccount.NewAccountFromValues(0, 1, 100, 0)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	amount, err := amount.NewAmountPositive(100)
	s.Require().NoError(err)

	s.Require().NoError(acc.BalanceReserve(amount))

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	s.Require().Equal(acc.Balance, acc.Balance)
}

func (s *Suite) TestSave_NotFound() {
	acc, err := domainAccount.NewAccountFromValues(1, 1, 100, 0)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().Error(err)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}
