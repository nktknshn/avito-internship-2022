package roles

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

var useCaseRoles = map[string][]auth.AuthUserRole{
	use_cases.GetBalance:        {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.Deposit:           {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.Reserve:           {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.ReserveCancel:     {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.ReserveConfirm:    {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.Transfer:          {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.AuthSignin:        {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.AuthSignup:        {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
	use_cases.AuthValidateToken: {auth.AuthUserRoleAccount, auth.AuthUserRoleAdmin},
}

func GetUseCaseAuthUserRoles(useCase string) []auth.AuthUserRole {
	if _, ok := useCaseRoles[useCase]; !ok {
		panic("use case " + useCase + " not found")
	}
	return useCaseRoles[useCase]
}
