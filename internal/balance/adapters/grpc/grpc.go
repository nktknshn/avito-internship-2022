package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/common/genproto/balance"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcAdapter struct {
	app app.Application
}

func New(application app.Application) *GrpcAdapter {
	return &GrpcAdapter{app: application}
}

func (g GrpcAdapter) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return auth.NewAuthInterceptor(
		g.app.AuthValidateToken,
		methodToRoles(),
	).Unary()
}

// func (g GrpcAdapter) Options() []grpc.ServerOption {
// 	return []grpc.ServerOption{
// 		grpc.ChainUnaryInterceptor(
// 			g.UnaryServerInterceptor(),
// 		),
// 	}
// }

func (g GrpcAdapter) GetBalance(ctx context.Context, request *balance.GetBalanceRequest) (*balance.GetBalanceResponse, error) {

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

func (g GrpcAdapter) Deposit(ctx context.Context, request *balance.DepositRequest) (*empty.Empty, error) {

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

func (g GrpcAdapter) Reserve(ctx context.Context, request *balance.ReserveRequest) (*empty.Empty, error) {

	in, err := reserve.NewInFromValues(
		request.UserId,
		request.ProductId,
		request.ProductTitle,
		request.OrderId,
		request.Amount,
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Reserve.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

func (g GrpcAdapter) ReserveCancel(ctx context.Context, request *balance.ReserveCancelRequest) (*empty.Empty, error) {

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

func (g GrpcAdapter) ReserveConfirm(ctx context.Context, request *balance.ReserveConfirmRequest) (*empty.Empty, error) {

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

func (g GrpcAdapter) Transfer(ctx context.Context, request *balance.TransferRequest) (*empty.Empty, error) {

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

func (g GrpcAdapter) AuthSignIn(ctx context.Context, request *balance.AuthSignInRequest) (*balance.AuthSignInResponse, error) {
	in, err := auth_signin.NewInFromValues(request.Username, request.Password)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := g.app.AuthSignin.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &balance.AuthSignInResponse{Token: out.Token.String()}, nil
}
