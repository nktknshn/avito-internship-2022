package accounts_pg

import (
	"context"
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
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
	s.ExecSQL("delete from accounts")
}

func (s *Suite) TestSave_Create_Success() {
	acc, err := domainAccount.NewAccountFromValues(
		0,
		fixtures.UserID_i64,
		fixtures.Amount100_i64,
		fixtures.Amount0_i64,
	)
	s.Require().NoError(err)

	_, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

}

func (s *Suite) TestSave_Update_Success() {
	acc, err := domainAccount.NewAccountFromValues(
		0,
		fixtures.UserID_i64,
		fixtures.Amount100_i64,
		fixtures.Amount0_i64,
	)

	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	amount, err := amount.NewPositive(100)
	s.Require().NoError(err)

	s.Require().NoError(acc.BalanceReserve(amount))

	acc2, err := s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	s.Require().Equal(acc.Balance, acc2.Balance)
}

func (s *Suite) TestSave_NotFound() {
	acc, err := domainAccount.NewAccountFromValues(
		1,
		fixtures.UserID_i64,
		fixtures.Amount100_i64,
		fixtures.Amount0_i64,
	)

	s.Require().NoError(err)

	_, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().Error(err)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}

func (s *Suite) TestGetByUserID_Success() {
	acc, err := domainAccount.NewAccountFromValues(
		0,
		fixtures.UserID_i64,
		fixtures.Amount100_i64,
		fixtures.Amount0_i64,
	)
	s.Require().NoError(err)

	_, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.GetByUserID(context.Background(), 1)
	s.Require().NoError(err)
	s.Require().Greater(acc.ID, domainAccount.AccountID(0))
	s.Require().Equal(acc.UserID, domain.UserID(1))
	s.Require().Equal(fixtures.Amount100, acc.Balance.GetAvailable())
	s.Require().Equal(fixtures.Amount0, acc.Balance.GetReserved())
}

func (s *Suite) TestGetByAccountID_Success() {
	acc, err := domainAccount.NewAccountFromValues(
		0,
		fixtures.UserID_i64,
		fixtures.Amount100_i64,
		fixtures.Amount0_i64,
	)

	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	acc2, err := s.accountsRepo.GetByAccountID(context.Background(), acc.ID)
	s.Require().NoError(err)
	s.Require().Equal(acc.ID, acc2.ID)
}

func (s *Suite) TestGetByAccountID_NotFound() {
	acc, err := s.accountsRepo.GetByAccountID(context.Background(), 1)
	s.Require().Error(err)
	s.Require().Nil(acc)
	s.Require().ErrorIs(err, domainAccount.ErrAccountNotFound)
}
