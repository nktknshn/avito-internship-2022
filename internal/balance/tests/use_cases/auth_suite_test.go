package use_cases_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/mocked"
)

func TestAuthSuiteTest(t *testing.T) {
	suite.Run(t, new(AuthSuiteTest))
}

type AuthSuiteTest struct {
	suite.Suite
	trm            *mocked.TrmManagerMock
	mockedAuthRepo *mocked.AuthRepositoryMock
	mockedHasher   *hasherVerifierMock
	mockedTokenGen *tokenManagerMock
	mockedTokenVal *tokenManagerMock
	// use cases with mocked dependencies

	mockedSignin   *auth_signin.AuthSigninUseCase
	mockedValidate *auth_validate_token.AuthValidateTokenUseCase
	mockedSignup   *auth_signup.AuthSignupUseCase
}

func (s *AuthSuiteTest) SetupTest() {
	s.trm = &mocked.TrmManagerMock{}
	s.mockedAuthRepo = &mocked.AuthRepositoryMock{}
	s.mockedHasher = &hasherVerifierMock{}
	s.mockedTokenGen = &tokenManagerMock{}
	s.mockedTokenVal = &tokenManagerMock{}

	// use cases with mocked dependencies
	s.mockedSignin = auth_signin.New(s.trm, s.mockedHasher, s.mockedTokenGen, s.mockedAuthRepo)
	s.mockedValidate = auth_validate_token.New(s.trm, s.mockedTokenVal, s.mockedAuthRepo)
	s.mockedSignup = auth_signup.New(s.trm, s.mockedHasher, s.mockedAuthRepo)
}
