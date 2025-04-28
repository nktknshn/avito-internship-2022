package grpc_test

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	authToken      string                = "token"
	userID_i64     int64                 = 1
	userID         domain.UserID         = 1
	authUserID_i64 int64                 = 2
	authUserID     domainAuth.AuthUserID = 1
	authIn                               = must.Must(auth_validate_token.NewInFromValues(authToken))
)
