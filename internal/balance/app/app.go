package app

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_list_users"
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
)

// Application это интерфейс приложения со всеми используемыми в приложении use case
type Application struct {
	// Auth
	AuthSignin        UseCase1[auth_signin.In, auth_signin.Out]
	AuthSignup        UseCase0[auth_signup.In]
	AuthValidateToken UseCase1[auth_validate_token.In, auth_validate_token.Out]
	AuthListUsers     UseCase1[auth_list_users.In, auth_list_users.Out]
	// Balance
	GetBalance          UseCase1[get_balance.In, get_balance.Out]
	Deposit             UseCase0[deposit.In]
	Reserve             UseCase0[reserve.In]
	ReserveCancel       UseCase0[reserve_cancel.In]
	ReserveConfirm      UseCase0[reserve_confirm.In]
	Transfer            UseCase0[transfer.In]
	ReportTransactions  UseCase1[report_transactions.In, report_transactions.Out]
	ReportRevenue       UseCase1[report_revenue.In, report_revenue.Out]
	ReportRevenueExport UseCase1[report_revenue_export.In, report_revenue_export.Out]
}
