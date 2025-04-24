package helpers

import (
	"github.com/avito-tech/go-transaction-manager/sql"
	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func GetTrm(s *testing_pg.TestSuitePg) *manager.Manager {
	trmFactory := trmsqlx.NewFactory(s.Conn, sql.NewSavePoint())
	trm, err := manager.New(trmFactory)
	if err != nil {
		panic(err)
	}
	return trm
}
