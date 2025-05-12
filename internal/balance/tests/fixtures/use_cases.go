package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_cancel"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/reserve_confirm"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	// Use cases

	InValidateToken = must.Must(auth_validate_token.NewInFromValues(AuthToken))
	InGetBalance    = must.Must(get_balance.NewInFromValues(UserID_i64))
	InDeposit100    = must.Must(deposit.NewInFromValues(
		UserID_i64,
		Amount100_i64,
		DepositSource_str,
	))
	InReserve100 = must.Must(reserve.NewInFromValues(
		UserID_i64,
		ProductID_i64,
		ProductTitle_str,
		OrderID_i64,
		Amount100_i64,
	))
	InReserveConfirm100 = must.Must(reserve_confirm.NewInFromValues(
		UserID_i64,
		ProductID_i64,
		OrderID_i64,
		Amount100_i64,
	))
	InReserveConfirm50 = must.Must(reserve_confirm.NewInFromValues(
		UserID_i64,
		ProductID_i64,
		OrderID_i64,
		Amount50_i64,
	))
	InReserveCancel100 = must.Must(reserve_cancel.NewInFromValues(
		UserID_i64,
		OrderID_i64,
		ProductID_i64,
		Amount100_i64,
	))
	InReserveCancel50 = must.Must(reserve_cancel.NewInFromValues(
		UserID_i64,
		OrderID_i64,
		ProductID_i64,
		Amount50_i64,
	))

	// Auth

	InAuthSignup = must.Must(auth_signup.NewInFromValues(
		UsernameAdmin_str,
		PasswordAdmin_str,
		AuthUserRole_str,
	))
	InAuthSignin = must.Must(auth_signin.NewInFromValues(
		UsernameAdmin_str,
		PasswordAdmin_str,
	))
	InAuthSigninInvalidPassword = must.Must(auth_signin.NewInFromValues(
		UsernameAdmin_str,
		PasswordAdmin_str+"invalid",
	))
)
