package auth_signin

import domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"

type In struct {
	Username domainAuth.AuthUserUsername
	Password domainAuth.AuthUserPassword
}

type Out struct {
	Token domainAuth.AuthUserToken
}
