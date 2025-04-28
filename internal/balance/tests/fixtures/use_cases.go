package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	AuthToken      string                = "token"
	AuthUserID_i64 int64                 = 1
	AuthUserID_str string                = "1"
	AuthUserID     domainAuth.AuthUserID = 1
	//
	UserID_i64 int64                 = 1
	UserID_str string                = "1"
	UserID     domainAuth.AuthUserID = 1
	//
	InValidateToken = must.Must(auth_validate_token.NewInFromValues(AuthToken))
	InGetBalance    = must.Must(get_balance.NewInFromValues(UserID_i64))
)
