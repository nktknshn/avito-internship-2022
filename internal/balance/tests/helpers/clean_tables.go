package helpers

import (
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func CleanTables(s *testing_pg.TestSuitePg) {
	s.ExecSql("DELETE FROM transactions_transfer")
	s.ExecSql("DELETE FROM transactions_spend")
	s.ExecSql("DELETE FROM transactions_deposit")
	s.ExecSql("DELETE FROM accounts")
	s.ExecSql("DELETE FROM auth_users")
	s.ExecSql("DELETE FROM auth_blacklist")
}
