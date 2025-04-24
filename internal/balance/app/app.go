package app

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/transfer"
	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
)

// Application is a struct that contains all the use cases
type Application struct {
	// Auth token
	AuthSignin        decorator.UseCaseQueryHandler[auth_signin.In, auth_signin.Out]
	AuthSignup        decorator.UseCaseCommandHandler[auth_signup.In]
	AuthValidateToken decorator.UseCaseQueryHandler[auth_validate_token.In, auth_validate_token.Out]
	//
	GetBalance     decorator.UseCaseQueryHandler[get_balance.In, get_balance.Out]
	Deposit        decorator.UseCaseCommandHandler[deposit.In]
	Reserve        decorator.UseCaseCommandHandler[reserve.In]
	ReserveCancel  decorator.UseCaseCommandHandler[reserve_cancel.In]
	ReserveConfirm decorator.UseCaseCommandHandler[reserve_confirm.In]
	Transfer       decorator.UseCaseCommandHandler[transfer.In]
}
