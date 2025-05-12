package helpers

import (
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func CleanTables(s *testing_pg.TestSuitePg) {
	s.ExecSQL("DELETE FROM transactions_transfer")
	s.ExecSQL("DELETE FROM transactions_spend")
	s.ExecSQL("DELETE FROM transactions_deposit")
	s.ExecSQL("DELETE FROM accounts")
	s.ExecSQL("DELETE FROM auth_users")
	s.ExecSQL("DELETE FROM auth_blacklist")
}
