package grpc

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
)

var grpcMethodToUseCase = map[string]string{
	"/balance.BalanceService/Deposit":           use_cases.Deposit,
	"/balance.BalanceService/GetBalance":        use_cases.GetBalance,
	"/balance.BalanceService/Reserve":           use_cases.Reserve,
	"/balance.BalanceService/ReserveCancel":     use_cases.ReserveCancel,
	"/balance.BalanceService/ReserveConfirm":    use_cases.ReserveConfirm,
	"/balance.BalanceService/Transfer":          use_cases.Transfer,
	"/balance.BalanceService/AuthSignin":        use_cases.AuthSignin,
	"/balance.BalanceService/AuthSignup":        use_cases.AuthSignup,
	"/balance.BalanceService/AuthValidateToken": use_cases.AuthValidateToken,
}
