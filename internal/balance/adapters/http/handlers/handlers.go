package handlers

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
)

type Handler interface {
	GetHandler() http.Handler
}

type Handlers struct {
	GetBalance     Handler
	Deposit        Handler
	Reserve        Handler
	ReserveCancel  Handler
	ReserveConfirm Handler
	Transfer       Handler
}

func CreateHandlers(app *app.Application) *Handlers {
	handlerGetBalance := get_balance.New(app.AuthValidateToken, app.GetBalance)
	handlerDeposit := deposit.New(app.AuthValidateToken, app.Deposit)
	handlerReserve := reserve.New(app.AuthValidateToken, app.Reserve)
	handlerReserveCancel := reserve_cancel.New(app.AuthValidateToken, app.ReserveCancel)
	handlerReserveConfirm := reserve_confirm.New(app.AuthValidateToken, app.ReserveConfirm)
	handlerTransfer := transfer.New(app.AuthValidateToken, app.Transfer)

	return &Handlers{
		GetBalance:     handlerGetBalance,
		Deposit:        handlerDeposit,
		Reserve:        handlerReserve,
		ReserveCancel:  handlerReserveCancel,
		ReserveConfirm: handlerReserveConfirm,
		Transfer:       handlerTransfer,
	}
}
