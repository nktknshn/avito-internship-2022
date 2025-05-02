package grpc_test

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
	"github.com/stretchr/testify/mock"
)

func (s *GrpcTestSuite) TestGetBalance_Success() {
	s.setupAuth(fixtures.AuthToken, auth_validate_token.Out{
		UserID: fixtures.AuthUserID,
		Role:   domainAuth.AuthUserRoleAdmin,
	}, nil)

	depositIn := must.Must(get_balance.NewInFromValues(fixtures.UserID_i64))

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, depositIn).Return(get_balance.Out{
		Available: fixtures.Amount100,
		Reserved:  fixtures.Amount0,
	}, nil)

	out, err := s.client.GetBalance(
		withAuthToken(s.T().Context(), fixtures.AuthToken),
		&balance.GetBalanceRequest{UserId: fixtures.UserID_i64},
	)

	s.Require().NoError(err)
	s.Require().Equal(int64(100), out.Available)
	s.Require().Equal(int64(0), out.Reserved)
}
