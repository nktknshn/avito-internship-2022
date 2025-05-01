package http

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
)

type Handler interface {
	GetHandler() http.Handler
}

type HttpAdapter struct {
	GetBalance         Handler
	Deposit            Handler
	Reserve            Handler
	ReserveCancel      Handler
	ReserveConfirm     Handler
	Transfer           Handler
	ReportTransactions Handler
	ReportRevenue      Handler
	AuthSignIn         Handler
}

func NewHttpAdapter(app *app.Application) *HttpAdapter {

	handlerGetBalance := get_balance.New(app.AuthValidateToken, app.GetBalance)
	handlerDeposit := deposit.New(app.AuthValidateToken, app.Deposit)
	handlerReserve := reserve.New(app.AuthValidateToken, app.Reserve)
	handlerReserveCancel := reserve_cancel.New(app.AuthValidateToken, app.ReserveCancel)
	handlerReserveConfirm := reserve_confirm.New(app.AuthValidateToken, app.ReserveConfirm)
	handlerTransfer := transfer.New(app.AuthValidateToken, app.Transfer)
	handlerSignIn := signin.New(app.AuthSignin)
	handlerReportTransactions := report_transactions.New(app.AuthValidateToken, app.ReportTransactions)
	handlerReportRevenue := report_revenue.New(app.AuthValidateToken, app.ReportRevenue)

	return &HttpAdapter{
		GetBalance:         handlerGetBalance,
		Deposit:            handlerDeposit,
		Reserve:            handlerReserve,
		ReserveCancel:      handlerReserveCancel,
		ReserveConfirm:     handlerReserveConfirm,
		Transfer:           handlerTransfer,
		AuthSignIn:         handlerSignIn,
		ReportTransactions: handlerReportTransactions,
		ReportRevenue:      handlerReportRevenue,
	}
}
