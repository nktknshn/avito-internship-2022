package server_grpc

import (
	"context"
	"net"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type RunningServer struct {
	GrpcServer *grpc.Server
	Listen     net.Listener
}

func RunGRPCServerOnAddr(addr string, logger logging.Logger, registerServer func(server *grpc.Server)) *RunningServer {

	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		// TODO
		// MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
		// Timeout:           s.cfg.Server.Timeout * time.Second,
		// MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
		// Time:              s.cfg.Server.Timeout * time.Minute,
	}))

	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal(context.Background(), "failed to listen: %v", "error", err)
	}

	go func() {
		err = grpcServer.Serve(listen)
		if err != nil {
			logger.Fatal(context.Background(), "failed to serve: %v", "error", err)
		}
	}()

	return &RunningServer{
		GrpcServer: grpcServer,
		Listen:     listen,
	}
}
