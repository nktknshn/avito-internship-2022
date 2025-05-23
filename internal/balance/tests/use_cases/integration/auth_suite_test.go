package use_cases_test

import (
	"testing"
	"time"

	trmsqlx "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/repositories/auth_pg"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/token_generator"
	"github.com/nktknshn/avito-internship-2022/pkg/password_hasher_argon"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/nktknshn/avito-internship-2022/pkg/token_generator_jwt"
)

func TestAuthUseCases(t *testing.T) {
	s := new(AuthSuiteIntegrationTest)
	s.SetPostgresMigrationsDir("../../../migrations/postgres")
	suite.Run(t, s)
}

type AuthSuiteIntegrationTest struct {
	testing_pg.TestSuitePg
	// real dependencies
	trm      trm.Manager
	authRepo *auth_pg.AuthRepository
	hasher   *password_hasher_argon.Hasher
	tokenGen token_generator.TokenGenerator[auth.AuthUserTokenClaims]
	tokenVal token_generator.TokenValidator[auth.AuthUserTokenClaims]
	// use cases
	signup   *auth_signup.AuthSignupUseCase
	signin   *auth_signin.AuthSigninUseCase
	validate *auth_validate_token.AuthValidateTokenUseCase
}

var (
	secretKey = []byte("secret")
)

func (s *AuthSuiteIntegrationTest) SetupTest() {
	s.trm = helpers.GetTrm(&s.TestSuitePg)
	s.authRepo = auth_pg.New(s.Conn, trmsqlx.DefaultCtxGetter)
	s.hasher = password_hasher_argon.New()

	s.tokenGen = token_generator_jwt.NewTokenGeneratorJWT[auth.AuthUserTokenClaims](secretKey, time.Hour*24)
	s.tokenVal = token_generator_jwt.NewTokenValidatorJWT[auth.AuthUserTokenClaims](secretKey)

	// real use cases
	s.signin = auth_signin.New(s.trm, s.hasher, s.tokenGen, s.authRepo)
	s.validate = auth_validate_token.New(s.trm, s.tokenVal, s.authRepo)
	s.signup = auth_signup.New(s.trm, s.hasher, s.authRepo)
}

func (s *AuthSuiteIntegrationTest) createAuthUser() {
	user, err := auth.NewAuthUserFromValues(
		0,
		fixtures.UsernameAdmin_str,
		fixtures.PasswordHashAdmin_str,
		fixtures.AuthUserRole_str,
	)
	s.Require().NoError(err)
	err = s.authRepo.CreateUser(s.Context(), user.Username, user.PasswordHash, user.Role)
	s.Require().NoError(err)
}
