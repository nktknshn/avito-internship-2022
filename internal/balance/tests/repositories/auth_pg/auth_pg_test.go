package auth_pg

import (
	"context"
	"errors"
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/auth_pg"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
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
	repo *auth_pg.AuthRepository
}

func (suite *TestSuiteAuthPg) SetupTest() {
	suite.repo = auth_pg.New(suite.Conn, trmsqlx.DefaultCtxGetter)
}

func (suite *TestSuiteAuthPg) TearDownTest() {
	suite.ExecSql("delete from auth_users")
}

func (s *TestSuiteAuthPg) TestGetUserByUsername_Success() {

	s.repo.CreateUser(context.Background(), "username", "password", domainAuth.AuthUserRoleAdmin)

	user, err := s.repo.GetUserByUsername(context.Background(), "username")
	s.NoError(err)
	s.Equal(domainAuth.AuthUserUsername("username"), user.Username)
	s.Equal(domainAuth.AuthUserPasswordHash("password"), user.PasswordHash)
	s.Equal(domainAuth.AuthUserRoleAdmin, user.Role)
}

func (s *TestSuiteAuthPg) TestGetUserByUsername_NotFound() {
	user, err := s.repo.GetUserByUsername(context.Background(), "test")
	s.Error(err)
	s.Nil(user)
	s.Require().True(errors.Is(err, domainAuth.ErrUserNotFound))
}

func (s *TestSuiteAuthPg) TestCreateAccount_Success() {
	err := s.repo.CreateUser(context.Background(), "username", "password", "role")
	s.NoError(err)
	rows, err := s.ExecSql("select * from auth_users")
	s.NoError(err)
	s.Equal(1, len(rows.Rows))
	s.Equal("username", rows.Rows[0]["username"])
	s.Equal("password", rows.Rows[0]["password_hash"])
	s.Equal("role", rows.Rows[0]["role"])
}

func (s *TestSuiteAuthPg) TestCreateAccount_DuplicateUsername() {
	err := s.repo.CreateUser(context.Background(), "test", "test", "test")
	s.NoError(err)
	err = s.repo.CreateUser(context.Background(), "test", "test", "test")
	s.Error(err)
	s.Require().True(errors.Is(err, domainAuth.ErrDuplicateUsername))
}
