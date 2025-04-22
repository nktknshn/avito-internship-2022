package auth_pg

import (
	"testing"

	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
)

func TestAuthPg(t *testing.T) {
	s := &TestSuiteAuthPg{}
	s.SetPostgresMigrationsDir("../../../migrations/postgres")
	suite.Run(t, s)
}

type TestSuiteAuthPg struct {
	testing_pg.TestSuitePg
}

func (suite *TestSuiteAuthPg) SetupTest() {

}

func (suite *TestSuiteAuthPg) TestCreateAccount() {

}
