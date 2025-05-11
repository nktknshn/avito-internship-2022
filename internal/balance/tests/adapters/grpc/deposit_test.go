package grpc_test

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type testCaseDeposit struct {
	name      string
	outErr    error
	errorCode codes.Code
	request   *balance.DepositRequest
}

func (s *GrpcTestSuite) TestDeposit() {
	testCases := []testCaseDeposit{
		{
			name: "success",
			request: &balance.DepositRequest{
				Amount: 100,
				Source: fixtures.DepositSource_str,
				UserId: fixtures.UserID_i64,
			},
			outErr:    nil,
			errorCode: codes.OK,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.setupAuthValid()

			s.app.DepositUseCaseMock.On("Handle", mock.Anything, mock.Anything).Return(tc.outErr)

			resp, err := s.client.Deposit(
				withAuthToken(s.T().Context(), fixtures.AuthToken),
				tc.request,
			)
			if tc.errorCode != codes.OK {
				s.Require().Error(err)
				s.Require().Nil(resp)
				s.Require().Equal(tc.errorCode, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(resp)
			}
		})
	}
}
