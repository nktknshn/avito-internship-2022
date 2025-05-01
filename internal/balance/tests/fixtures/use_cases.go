package fixtures

import (
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
)
