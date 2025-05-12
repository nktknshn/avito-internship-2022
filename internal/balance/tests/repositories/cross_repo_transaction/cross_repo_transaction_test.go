package cross_repo_transaction

import (
	"context"
	"errors"
	"testing"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/accounts_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/transactions_pg"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
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
	s.ExecSQL("delete from accounts")
	s.ExecSQL("delete from transactions_deposit")
	s.ExecSQL("delete from transactions_spend")
	s.ExecSQL("delete from transactions_transfer")
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
			uuid.Nil,
			acc.ID.Value(),
			fixtures.UserID_i64,
			fixtures.OrderID_i64,
			fixtures.ProductID_i64,
			fixtures.ProductTitle_str,
			fixtures.AmountPositive_i64,
			domainTransaction.TransactionSpendStatusReserved.Value(),
			time.Now(), time.Now(),
		)
		s.Require().NoError(err)
		_, err = s.transactionsRepo.SaveTransactionSpend(ctx, ts)
		s.Require().NoError(err)
		return errors.New("test error")
	})

	s.Require().Error(err)

	rows, err := s.ExecSQL("select * from transactions_spend")
	s.Require().NoError(err)
	s.Require().Empty(rows.Rows)

	rows, err = s.ExecSQL("select * from accounts")
	s.Require().NoError(err)
	s.Require().Empty(rows.Rows)
}
