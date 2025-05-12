package grpc_test

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app_impl"
	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/grpc/server"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/e2e"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/helpers"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

func TestE2ESuite(t *testing.T) {
	s := &E2ESuite{}
	s.NeedsPostgres = true
	s.MigrationsDir = "../../../migrations/postgres"
	suite.Run(t, s)
}

type E2ESuite struct {
	testing_pg.TestSuitePg
	server     *server.BalanceGrpcServer
	clientConn *grpc.ClientConn
	client     balance.BalanceServiceClient
	cleanup    func()
}

func (s *E2ESuite) SetupSubTest() {

}

func (s *E2ESuite) TearDownSubTest() {

}

func (s *E2ESuite) SetupTest() {
	cfg, err := config.LoadConfigFromReader(strings.NewReader(e2e.Cfg))
	s.Require().NoError(err)

	port, err := s.DT.GetRunningPort()
	s.Require().NoError(err)

	cfg.Postgres.Addr = "localhost:" + port
	cfg.Postgres.User = "postgres"
	cfg.Postgres.Database = "postgres"
	cfg.Postgres.Schema = "public"
	cfg.Postgres.Password = s.DT.GetDockerConfig().DockerPassword
	cfg.Postgres.MigrationsDir = ""

	app, cleanup, err := app_impl.NewApplication(context.Background(), cfg)
	s.Require().NoError(err)
	s.cleanup = cleanup
	s.server = server.NewGrpcServer(cfg, app)

	s.Require().NoError(s.server.Init(context.Background()))
	s.Require().NoError(s.server.Run(context.Background()))

	insec := grpc.WithTransportCredentials(insecure.NewCredentials())
	client, err := grpc.NewClient(s.server.GetConfig().GRPC.Addr, insec)
	s.Require().NoError(err)
	s.clientConn = client
	s.client = balance.NewBalanceServiceClient(client)
}

func (s *E2ESuite) TearDownTest() {
	s.Require().NoError(s.server.Shutdown(context.Background()))
	s.cleanup()
	helpers.CleanTables(&s.TestSuitePg)
}

func (s *E2ESuite) TestGrpcServer() {
	s.Run("should run grpc server", func() {
		_, err := s.client.GetBalance(context.Background(), &balance.GetBalanceRequest{
			UserId: 1,
		})
		s.Require().Error(err)
		s.Require().ErrorIs(err, status.Error(codes.Unauthenticated, "missing auth token"))
	})
}
