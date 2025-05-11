package grpc_test

import (
	"errors"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type testCaseDeposit struct {
	desc      string
	request   *balance.GetBalanceRequest
	out       get_balance.Out
	outErr    error
	errorCode codes.Code
}

func (s *GrpcTestSuite) TestGetBalance() {
	testCases := []testCaseDeposit{
		{
			desc:    "success",
			request: &balance.GetBalanceRequest{UserId: fixtures.UserID_i64},
			out: get_balance.Out{
				Available: fixtures.Amount100,
				Reserved:  fixtures.Amount0,
			},
			outErr:    nil,
			errorCode: codes.OK,
		},
		{
			desc:      "error",
			request:   &balance.GetBalanceRequest{UserId: fixtures.UserID_i64},
			out:       get_balance.Out{},
			outErr:    errors.New("error"),
			errorCode: codes.Internal,
		},
		{
			desc:      "invalid request",
			request:   &balance.GetBalanceRequest{UserId: fixtures.UserID_i64_invalid},
			out:       get_balance.Out{},
			outErr:    errors.New("error"),
			errorCode: codes.InvalidArgument,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.desc, func() {
			s.setupAuthValid()

			s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, fixtures.InGetBalance).Return(tc.out, tc.outErr)

			resp, err := s.client.GetBalance(
				withAuthToken(s.T().Context(), fixtures.AuthToken),
				tc.request,
			)

			if tc.errorCode != codes.OK {
				s.Require().Error(err)
				s.Require().Nil(resp)
				s.Require().Equal(tc.errorCode, status.Code(err))
			} else {
				s.Require().NoError(err)
				s.Require().Equal(tc.out.Available.Value(), resp.Available)
				s.Require().Equal(tc.out.Reserved.Value(), resp.Reserved)
			}
		})
	}
}
