package helpers

import (
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func CleanTables(s *testing_pg.TestSuitePg) {
	s.ExecSQLMust("DELETE FROM transactions_transfer")
	s.ExecSQLMust("DELETE FROM transactions_spend")
	s.ExecSQLMust("DELETE FROM transactions_deposit")
	s.ExecSQLMust("DELETE FROM accounts")
	s.ExecSQLMust("DELETE FROM auth_users")
	s.ExecSQLMust("DELETE FROM auth_blacklist")
}
