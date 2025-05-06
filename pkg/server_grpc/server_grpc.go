package server_grpc

import (
	"net"

	"github.com/nktknshn/avito-internship-2022/internal/common/logging"
	"google.golang.org/grpc"
)

type RunningServer struct {
	GrpcServer *grpc.Server
	Listen     net.Listener
}

func RunGRPCServerOnAddr(addr string, logger logging.Logger, registerServer func(server *grpc.Server), opts ...grpc.ServerOption) *RunningServer {

	grpcServer := grpc.NewServer(opts...)

	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatal("failed to listen: %v", "error", err)
	}

	go func() {
		err = grpcServer.Serve(listen)
		if err != nil {
			logger.Fatal("failed to serve: %v", "error", err)
		}
	}()

	return &RunningServer{
		GrpcServer: grpcServer,
		Listen:     listen,
	}
}
