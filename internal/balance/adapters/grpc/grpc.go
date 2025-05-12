package grpc

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/grpc/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
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
	app *app.Application
}

func New(application *app.Application) *GrpcAdapter {
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

	in, err := get_balance.NewInFromValues(request.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := g.app.GetBalance.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &balance.GetBalanceResponse{Available: out.Available.Value(), Reserved: out.Reserved.Value()}, nil
}

func (g GrpcAdapter) Deposit(ctx context.Context, request *balance.DepositRequest) (*emptypb.Empty, error) {

	in, err := deposit.NewInFromValues(request.GetUserId(), request.GetAmount(), request.GetSource())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Deposit.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcAdapter) Reserve(ctx context.Context, request *balance.ReserveRequest) (*emptypb.Empty, error) {

	in, err := reserve.NewInFromValues(
		request.GetUserId(),
		request.GetProductId(),
		request.GetProductTitle(),
		request.GetOrderId(),
		request.GetAmount(),
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Reserve.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcAdapter) ReserveCancel(ctx context.Context, request *balance.ReserveCancelRequest) (*emptypb.Empty, error) {

	in, err := reserve_cancel.NewInFromValues(request.GetUserId(), request.GetOrderId(), request.GetProductId(), request.GetAmount())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.ReserveCancel.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcAdapter) ReserveConfirm(ctx context.Context, request *balance.ReserveConfirmRequest) (*emptypb.Empty, error) {

	in, err := reserve_confirm.NewInFromValues(request.GetUserId(), request.GetOrderId(), request.GetProductId(), request.GetAmount())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.ReserveConfirm.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcAdapter) Transfer(ctx context.Context, request *balance.TransferRequest) (*emptypb.Empty, error) {

	in, err := transfer.NewInFromValues(request.GetFrom(), request.GetTo(), request.GetAmount())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = g.app.Transfer.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g GrpcAdapter) ReportTransactions(
	ctx context.Context,
	request *balance.ReportTransactionsRequest,
) (*balance.ReportTransactionsResponse, error) {

	in, err := report_transactions.NewInFromValues(
		request.GetUserId(),
		request.GetCursor(),
		request.GetLimit(),
		request.GetSorting(),
		request.GetSortingDirection(),
	)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := g.app.ReportTransactions.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	transactions := make([]*balance.ReportTransactionsTransaction, len(out.Transactions))

	for i, transaction := range out.Transactions {
		switch t := transaction.(type) {
		case *report_transactions.OutTransactionDeposit:
			transactions[i] = &balance.ReportTransactionsTransaction{
				Transaction: &balance.ReportTransactionsTransaction_Deposit{
					Deposit: &balance.ReportTransactionsTransactionDeposit{
						Id:        t.ID.Value().String(),
						Amount:    t.Amount.Value(),
						Source:    t.Source.Value(),
						Status:    t.Status.Value(),
						CreatedAt: t.CreatedAt.Format(time.RFC3339),
						UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
					},
				},
			}
		case *report_transactions.OutTransactionSpend:
			transactions[i] = &balance.ReportTransactionsTransaction{
				Transaction: &balance.ReportTransactionsTransaction_Spend{
					Spend: &balance.ReportTransactionsTransactionSpend{
						Id:           t.ID.Value().String(),
						AccountId:    t.AccountID.Value(),
						OrderId:      t.OrderID.Value(),
						ProductId:    t.ProductID.Value(),
						ProductTitle: t.ProductTitle.Value(),
						Amount:       t.Amount.Value(),
						Status:       t.Status.Value(),
						CreatedAt:    t.CreatedAt.Format(time.RFC3339),
						UpdatedAt:    t.UpdatedAt.Format(time.RFC3339),
					},
				},
			}
		case *report_transactions.OutTransactionTransfer:
			transactions[i] = &balance.ReportTransactionsTransaction{
				Transaction: &balance.ReportTransactionsTransaction_Transfer{
					Transfer: &balance.ReportTransactionsTransactionTransfer{
						Id:        t.ID.Value().String(),
						From:      t.From.Value(),
						To:        t.To.Value(),
						Amount:    t.Amount.Value(),
						Status:    t.Status.Value(),
						CreatedAt: t.CreatedAt.Format(time.RFC3339),
						UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
					},
				},
			}
		default:
			return nil, status.Error(codes.Internal, "unknown transaction type")
		}
	}

	return &balance.ReportTransactionsResponse{
		Transactions: transactions,
		Cursor:       string(out.Cursor),
		HasMore:      out.HasMore,
	}, nil
}

func (g GrpcAdapter) ReportRevenue(ctx context.Context, request *balance.ReportRevenueRequest) (*balance.ReportRevenueResponse, error) {

	in, err := report_revenue.NewInFromValues(int(request.GetYear()), int(request.GetMonth()))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := g.app.ReportRevenue.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	records := make([]*balance.ReportRevenueRecord, len(out.Records))

	for i, record := range out.Records {
		records[i] = &balance.ReportRevenueRecord{
			ProductId:    record.ProductID.Value(),
			ProductTitle: record.ProductTitle.Value(),
			TotalRevenue: record.TotalRevenue,
		}
	}

	return &balance.ReportRevenueResponse{
		Records: records,
	}, nil
}

func (g GrpcAdapter) AuthSignIn(ctx context.Context, request *balance.AuthSignInRequest) (*balance.AuthSignInResponse, error) {
	in, err := auth_signin.NewInFromValues(request.GetUsername(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := g.app.AuthSignin.Handle(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &balance.AuthSignInResponse{Token: out.Token.String()}, nil
}
