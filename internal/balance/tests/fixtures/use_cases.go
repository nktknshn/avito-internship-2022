package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signin"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_signup"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/deposit"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	// Use cases
	InValidateToken = must.Must(auth_validate_token.NewInFromValues(AuthToken))
	InGetBalance    = must.Must(get_balance.NewInFromValues(UserID_i64))
	InDeposit       = must.Must(deposit.NewInFromValues(UserID_i64, Amount100_i64, DepositSource_str))
	InAuthSignup    = must.Must(auth_signup.NewInFromValues(
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
