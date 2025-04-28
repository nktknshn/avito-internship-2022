package transactions_pg_test

import (
	"context"
	"testing"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/google/uuid"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
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

func (s *Suite) TestSaveTransactionDeposit_Success() {
	acc, err := domainAccount.NewAccount(1)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	transaction := must.Must(domainTransaction.NewTransactionDepositFromValues(
		uuid.New(), acc.ID.Value(), 1, "test", "confirmed", 100, time.Now(), time.Now(),
	))

	transaction, err = s.transactionsRepo.SaveTransactionDeposit(context.Background(), transaction)
	s.Require().NoError(err)

	s.Require().Greater(transaction.ID, domainTransaction.TransactionDepositID(uuid.Nil))

	rows, err := s.ExecSql("select * from transactions_deposit")
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))
	s.Require().Equal(transaction.ID.Value(), rows.Rows[0]["id"])
	s.Require().Equal(acc.ID.Value(), rows.Rows[0]["account_id"])
	s.Require().Equal(int64(1), rows.Rows[0]["user_id"])
	s.Require().Equal("test", rows.Rows[0]["deposit_source"])
	s.Require().Equal("confirmed", rows.Rows[0]["status"])
	s.Require().Equal(int64(100), rows.Rows[0]["amount"])
}
