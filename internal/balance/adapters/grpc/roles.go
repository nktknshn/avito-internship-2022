package grpc

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/roles"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

const balanceServicePath = "/balance.BalanceService/"

func init() {
	// проверить, что все методы GrpcServer соответствуют use cases
	server := GrpcAdapter{}
	typ := reflect.TypeOf(server)
	var ignoredMethods = []string{
		"UnaryServerInterceptor",
		"Options",
	}
	for i := range typ.NumMethod() {
		method := typ.Method(i)
		if slices.Contains(ignoredMethods, method.Name) {
			continue
		}
		_, ok := roles.GetUseCaseAuthUserRoles(method.Name)
		if !ok {
			panic(fmt.Sprintf("method %s has no use case", method.Name))
		}
	}
}

func methodToRoles() map[string][]domainAuth.AuthUserRole {
	rolesMap := map[string][]domainAuth.AuthUserRole{}
	for _, uc := range roles.GetUseCases() {
		rolesMap[balanceServicePath+uc] = roles.GetUseCaseAuthUserRolesMust(uc)
	}
	return rolesMap
}
