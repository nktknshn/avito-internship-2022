package auth_validate_token

import domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"

type In struct {
	Token string
}

type Out struct {
	UserID domainAuth.AuthUserID
	Role   domainAuth.AuthUserRole
}
