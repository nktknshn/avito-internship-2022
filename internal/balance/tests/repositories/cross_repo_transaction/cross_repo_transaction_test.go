package cross_repo_transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/google/uuid"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
)

func TestCrossRepoTransaction(t *testing.T) {
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
	s.accountsRepo = accounts_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
	s.transactionsRepo = transactions_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
}

func (s *Suite) TearDownTest() {
	s.ExecSql("delete from accounts")
	s.ExecSql("delete from transactions_deposit")
	s.ExecSql("delete from transactions_spend")
	s.ExecSql("delete from transactions_transfer")
}

// TestCrossRepoTransaction_Fail проверяет, что транзакция работает в рамках двух репозиториев
func (s *Suite) TestCrossRepoTransaction_Fail() {
	acc, err := domainAccount.NewAccountFromValues(0, 1, 0, 0)
	s.Require().NoError(err)

	trm := helpers.GetTrm(&s.TestSuitePg)

	err = trm.Do(s.Context(), func(ctx context.Context) error {
		acc, err = s.accountsRepo.Save(ctx, acc)
		s.Require().NoError(err)
		ts, err := domainTransaction.NewTransactionSpendFromValues(
			uuid.Nil, acc.ID.Value(), 1, 1, 1, 100, domainTransaction.TransactionSpendStatusReserved.Value(), time.Now(), time.Now(),
		)
		s.Require().NoError(err)
		_, err = s.transactionsRepo.SaveTransactionSpend(ctx, ts)
		s.Require().NoError(err)
		return errors.New("test error")
	})

	s.Require().Error(err)

	rows, err := s.ExecSql("select * from transactions_spend")
	s.Require().NoError(err)
	s.Require().Equal(0, len(rows.Rows))

	rows, err = s.ExecSql("select * from accounts")
	s.Require().NoError(err)
	s.Require().Equal(0, len(rows.Rows))
}
