package grpc_test

import (
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *GrpcTestSuite) TestGetBalance_Unauthenticated() {
	s.Run("missing auth token", func() {
		_, err := s.client.GetBalance(
			s.T().Context(),
			&balance.GetBalanceRequest{UserId: 1},
		)

		s.Require().Equal(status.Code(err), codes.Unauthenticated)
	})

	s.Run("invalid auth token", func() {
		s.setupAuth(fixtures.AuthToken, auth_validate_token.Out{}, errors.New("unauthorized"))

		_, err := s.client.GetBalance(
			withAuthToken(s.T().Context(), fixtures.AuthToken),
			&balance.GetBalanceRequest{UserId: 1},
		)

		s.Require().Equal(status.Code(err), codes.Unauthenticated)
	})

	s.Run("invalid role", func() {
		s.setupAuth(fixtures.AuthToken, auth_validate_token.Out{
			UserID: fixtures.AuthUserID,
			Role:   domainAuth.AuthUserRoleReport,
		}, nil)

		_, err := s.client.GetBalance(
			withAuthToken(s.T().Context(), fixtures.AuthToken),
			&balance.GetBalanceRequest{UserId: 1},
		)

		s.Require().Equal(status.Code(err), codes.PermissionDenied)
	})
}
