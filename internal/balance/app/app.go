package app

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue_export"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
)

// Application это интерфейс приложения со всеми используемыми в приложении use case
type Application struct {
	// Auth
	AuthSignin        decorator.UseCase1Handler[auth_signin.In, auth_signin.Out]
	AuthSignup        decorator.UseCase0Handler[auth_signup.In]
	AuthValidateToken decorator.UseCase1Handler[auth_validate_token.In, auth_validate_token.Out]
	// Balance
	GetBalance          decorator.UseCase1Handler[get_balance.In, get_balance.Out]
	Deposit             decorator.UseCase0Handler[deposit.In]
	Reserve             decorator.UseCase0Handler[reserve.In]
	ReserveCancel       decorator.UseCase0Handler[reserve_cancel.In]
	ReserveConfirm      decorator.UseCase0Handler[reserve_confirm.In]
	Transfer            decorator.UseCase0Handler[transfer.In]
	ReportTransactions  decorator.UseCase1Handler[report_transactions.In, report_transactions.Out]
	ReportRevenue       decorator.UseCase1Handler[report_revenue.In, report_revenue.Out]
	ReportRevenueExport decorator.UseCase1Handler[report_revenue_export.In, report_revenue_export.Out]
}
