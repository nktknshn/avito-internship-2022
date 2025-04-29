package transactions_pg_test

import (
	"context"
	"time"

	"github.com/google/uuid"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

func (s *Suite) TestSaveTransactionDeposit_Success() {
	acc, err := domainAccount.NewAccount(1)
	s.Require().NoError(err)

	acc, err = s.accountsRepo.Save(context.Background(), acc)
	s.Require().NoError(err)

	transaction := must.Must(domainTransaction.NewTransactionDepositFromValues(
		uuid.New(),
		acc.ID.Value(),
		fixtures.UserID_i64,
		fixtures.DepositSource_str,
		domainTransaction.TransactionDepositStatusConfirmed.Value(),
		fixtures.AmountPositive100_i64,
		time.Now(),
		time.Now(),
	))

	transaction, err = s.transactionsRepo.SaveTransactionDeposit(context.Background(), transaction)
	s.Require().NoError(err)

	s.Require().NotEqual(transaction.ID, domainTransaction.TransactionDepositID(uuid.Nil))

	rows, err := s.ExecSql("select * from transactions_deposit")
	s.Require().NoError(err)
	s.Require().Equal(1, len(rows.Rows))
	s.Require().Equal(transaction.ID.Value().String(), rows.Rows[0]["id"])
	s.Require().Equal(acc.ID.Value(), rows.Rows[0]["account_id"])
	s.Require().Equal(fixtures.UserID_i64, rows.Rows[0]["user_id"])
	s.Require().Equal(fixtures.DepositSource_str, rows.Rows[0]["deposit_source"])
	s.Require().Equal(domainTransaction.TransactionDepositStatusConfirmed.Value(), rows.Rows[0]["status"])
	s.Require().Equal(fixtures.AmountPositive100_i64, rows.Rows[0]["amount"])
}
