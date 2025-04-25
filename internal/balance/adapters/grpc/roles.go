package grpc

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/roles"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

var AccessibleRoles map[string][]domainAuth.AuthUserRole

func init() {
	AccessibleRoles = make(map[string][]domainAuth.AuthUserRole)
	for method, uc := range grpcMethodToUseCase {
		AccessibleRoles[method] = roles.GetUseCaseAuthUserRoles(uc)
	}
}
