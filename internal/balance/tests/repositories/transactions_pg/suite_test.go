package transactions_pg_test

import (
	"context"
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func TestTransactionsPg(t *testing.T) {
	s := &Suite{}
	s.SetPostgresMigrationsDir("../../../migrations/postgres")
	suite.Run(t, s)
}

type Suite struct {
	testing_pg.TestSuitePg
	accountsRepo     *accounts_pg.AccountsRepository
	transactionsRepo *transactions_pg.TransactionsRepository
}

func (s *Suite) SetupTest() {
	s.transactionsRepo = transactions_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
	s.accountsRepo = accounts_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
}

func (s *Suite) TearDownTest() {
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *Suite) getAccount1() *domainAccount.Account {
	acc, err := s.accountsRepo.Save(context.Background(), &domainAccount.Account{
		UserID: fixtures.UserID,
	})
	s.Require().NoError(err)
	s.Require().NotNil(acc)
	return acc
}

func (s *Suite) getAccount2() *domainAccount.Account {
	acc, err := s.accountsRepo.Save(context.Background(), &domainAccount.Account{
		UserID: fixtures.UserID_2,
	})
	s.Require().NoError(err)
	s.Require().NotNil(acc)
	return acc
}
