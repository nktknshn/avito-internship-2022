package grpc_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	adaptersGrpc "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/mocked"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
)

func TestGrpcTestSuite(t *testing.T) {
	suite.Run(t, new(GrpcTestSuite))
}

type GrpcTestSuite struct {
	suite.Suite
	app         *mocked.AppMocked
	grpcAdapter *adaptersGrpc.GrpcAdapter
	grpcServer  *grpc.Server
	listen      net.Listener
	clientConn  *grpc.ClientConn
	client      balance.BalanceServiceClient
}

func (s *GrpcTestSuite) SetupSubTest() {
	s.SetupTest()
}

func (s *GrpcTestSuite) TearDownSubTest() {
	s.TearDownTest()
}

func (s *GrpcTestSuite) SetupTest() {
	s.app = mocked.NewMockedApp()
	s.grpcAdapter = adaptersGrpc.New(&s.app.Application)
	options := []grpc.ServerOption{
		grpc.UnaryInterceptor(s.grpcAdapter.UnaryServerInterceptor()),
	}
	s.grpcServer = grpc.NewServer(options...)
	balance.RegisterBalanceServiceServer(s.grpcServer, s.grpcAdapter)

	listen, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)
	s.listen = listen

	go func() {
		serveErr := s.grpcServer.Serve(s.listen)
		if serveErr != nil {
			panic(serveErr)
		}
	}()

	insec := grpc.WithTransportCredentials(insecure.NewCredentials())
	client, err := grpc.NewClient(s.listen.Addr().String(), insec)
	s.Require().NoError(err)
	s.clientConn = client
	s.client = balance.NewBalanceServiceClient(client)
}

func (s *GrpcTestSuite) TearDownTest() {
	s.clientConn.Close()
	s.grpcServer.Stop()
	s.listen.Close()
}

func (s *GrpcTestSuite) setupAuth(token string, returnOut auth_validate_token.Out, returnErr error) {
	s.app.SetupAuth(token, returnOut, returnErr)
}

func (s *GrpcTestSuite) setupAuthValid() {
	s.setupAuth(fixtures.AuthToken, auth_validate_token.Out{
		UserID: fixtures.AuthUserID,
		Role:   domainAuth.AuthUserRoleAdmin,
	}, nil)
}
