package auth_signin

import (
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
)

type Out struct {
	Token domainAuth.AuthUserToken
}
