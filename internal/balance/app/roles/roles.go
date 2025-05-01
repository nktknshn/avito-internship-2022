package roles

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

// пустой слайс означает, что метод доступен без авторизации
var useCaseRoles = map[string][]auth.AuthUserRole{
	use_cases.NameGetBalance:         {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.NameDeposit:            {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.NameReserve:            {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.NameReserveCancel:      {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.NameReserveConfirm:     {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.NameTransfer:           {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.NameReportTransactions: {auth.AuthUserRoleReport, auth.AuthUserRoleAdmin},
	use_cases.NameAuthSignin:         {},
}

func GetUseCaseAuthUserRoles(useCase string) ([]auth.AuthUserRole, bool) {
	roles, ok := useCaseRoles[useCase]
	return roles, ok
}

func GetUseCaseAuthUserRolesMust(useCase string) []auth.AuthUserRole {
	roles, ok := GetUseCaseAuthUserRoles(useCase)
	if !ok {
		panic("use case " + useCase + " not found")
	}
	return roles
}

func GetUseCases() []string {
	useCases := []string{}
	for useCase := range useCaseRoles {
		useCases = append(useCases, useCase)
	}
	return useCases
}
