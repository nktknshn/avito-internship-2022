package auth_pg

import (
	"testing"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"

	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/auth_pg"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
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

func (s *TestSuiteAuthPg) SetupTest() {
	s.repo = auth_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
}

func (s *TestSuiteAuthPg) TearDownTest() {
	s.ExecSQL("delete from auth_users")
}

func (s *TestSuiteAuthPg) TestGetUserByUsername_Success() {

	s.repo.CreateUser(s.Context(), "username", "password", domainAuth.AuthUserRoleAdmin)

	user, err := s.repo.GetUserByUsername(s.Context(), "username")
	s.Require().NoError(err)
	s.Require().Equal(domainAuth.AuthUserUsername("username"), user.Username)
	s.Require().Equal(domainAuth.AuthUserPasswordHash("password"), user.PasswordHash)
	s.Require().Equal(domainAuth.AuthUserRoleAdmin, user.Role)
}

func (s *TestSuiteAuthPg) TestGetUserByUsername_NotFound() {
	user, err := s.repo.GetUserByUsername(s.Context(), "test")
	s.Require().Error(err)
	s.Require().Nil(user)
	s.Require().ErrorIs(err, domainAuth.ErrAuthUserNotFound)
}

func (s *TestSuiteAuthPg) TestCreateUser_Success() {
	err := s.repo.CreateUser(s.Context(), "username", "password", "role")
	s.Require().NoError(err)
	rows, err := s.ExecSQL("select * from auth_users")
	s.Require().NoError(err)
	s.Require().Len(rows.Rows, 1)
	s.Require().Equal("username", rows.Rows[0]["username"])
	s.Require().Equal("password", rows.Rows[0]["password_hash"])
	s.Require().Equal("role", rows.Rows[0]["role"])
}

func (s *TestSuiteAuthPg) TestCreateUser_DuplicateUsername() {
	err := s.repo.CreateUser(s.Context(), "test", "test", "test")
	s.Require().NoError(err)
	err = s.repo.CreateUser(s.Context(), "test", "test", "test")
	s.Require().Error(err)
	s.Require().ErrorIs(err, domainAuth.ErrDuplicateUsername)
}

func (s *TestSuiteAuthPg) TestListUsers_Empty() {
	users, err := s.repo.ListUsers(s.Context())
	s.Require().NoError(err)
	s.Require().Empty(users)
}

func (s *TestSuiteAuthPg) TestListUsers_Success() {
	err := s.repo.CreateUser(s.Context(), "username", "password", "admin")
	s.Require().NoError(err)

	err = s.repo.CreateUser(s.Context(), "username2", "password2", "admin")
	s.Require().NoError(err)

	users, err := s.repo.ListUsers(s.Context())
	s.Require().NoError(err)
	s.Require().Len(users, 2)
	s.Require().Equal("username", users[0].Username.Value())
	s.Require().Equal("username2", users[1].Username.Value())
}
