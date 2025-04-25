package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	app app.Application
}

func NewGrpcServer(application app.Application) GrpcServer {
	return GrpcServer{app: application}
}

func (g GrpcServer) GetBalance(ctx context.Context, request *balance.GetBalanceRequest) (*balance.GetBalanceResponse, error) {

	in, err := get_balance.NewInFromValues(request.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := g.app.GetBalance.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &balance.GetBalanceResponse{Available: out.Available, Reserved: out.Reserved}, nil
}

func (g GrpcServer) Deposit(ctx context.Context, request *balance.DepositRequest) (*empty.Empty, error) {

	in, err := deposit.NewInFromValues(request.UserId, request.Amount, request.Source)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Deposit.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) Reserve(ctx context.Context, request *balance.ReserveRequest) (*empty.Empty, error) {

	in, err := reserve.NewInFromValues(request.UserId, request.ProductId, request.OrderId, request.Amount)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Reserve.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) ReserveCancel(ctx context.Context, request *balance.ReserveCancelRequest) (*empty.Empty, error) {

	in, err := reserve_cancel.NewInFromValues(request.UserId, request.OrderId, request.ProductId, request.Amount)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.ReserveCancel.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) ReserveConfirm(ctx context.Context, request *balance.ReserveConfirmRequest) (*empty.Empty, error) {

	in, err := reserve_confirm.NewInFromValues(request.UserId, request.OrderId, request.ProductId, request.Amount)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.ReserveConfirm.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcServer) Transfer(ctx context.Context, request *balance.TransferRequest) (*empty.Empty, error) {

	in, err := transfer.NewInFromValues(request.From, request.To, request.Amount)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Transfer.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
