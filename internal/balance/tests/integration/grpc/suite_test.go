package grpc_test

import (
	"context"
	"net"
	"testing"

	adaptersGrpc "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

func TestGrpcTestSuite(t *testing.T) {
	suite.Run(t, new(GrpcTestSuite))
}

type GrpcTestSuite struct {
	suite.Suite
	app         *helpers.AppMocked
	grpcAdapter *adaptersGrpc.GrpcAdapter
	grpcServer  *grpc.Server
	listen      net.Listener
}

func (s *GrpcTestSuite) SetupTest() {
	s.app = helpers.NewMockedApp()
	s.grpcAdapter = adaptersGrpc.New(s.app.Application)
	options := []grpc.ServerOption{
		grpc.UnaryInterceptor(s.grpcAdapter.UnaryServerInterceptor()),
	}
	options = append(options, s.grpcAdapter.Options()...)
	s.grpcServer = grpc.NewServer(options...)
	balance.RegisterBalanceServiceServer(s.grpcServer, s.grpcAdapter)

	listen, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)
	s.listen = listen
	go func() {
		s.grpcServer.Serve(s.listen)
	}()
}

func (s *GrpcTestSuite) TearDownTest() {
	s.grpcServer.Stop()
	s.listen.Close()
}

func (s *GrpcTestSuite) client() (balance.BalanceServiceClient, func()) {
	client, err := grpc.NewClient(s.listen.Addr().String(), grpc.WithInsecure())
	s.Require().NoError(err)
	return balance.NewBalanceServiceClient(client), func() {
		client.Close()
	}
}

func (s *GrpcTestSuite) TestGetBalance() {
	in, err := get_balance.NewInFromValues(1)
	s.Require().NoError(err)

	s.app.GetBalanceUseCaseMock.On("Handle", mock.Anything, in).Return(get_balance.Out{}).Return(get_balance.Out{
		Available: 100,
		Reserved:  0,
	}, nil)

	client, closeClient := s.client()
	defer closeClient()

	out, err := client.GetBalance(context.Background(), &balance.GetBalanceRequest{
		UserId: 1,
	})

	s.Require().NoError(err)
	s.Require().Equal(100, out.Available)
	s.Require().Equal(0, out.Reserved)
}
