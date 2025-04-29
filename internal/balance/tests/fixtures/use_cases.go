package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	AuthToken      string                = "token"
	AuthUserID_i64 int64                 = 1
	AuthUserID_str string                = "1"
	AuthUserID     domainAuth.AuthUserID = 1
	//
	UserID_i64 int64         = 1
	UserID_str string        = "1"
	UserID     domain.UserID = 1
	//
	UserID_2_i64 int64         = 2
	UserID_2_str string        = "2"
	UserID_2     domain.UserID = 2
	//
	UserID_3_i64 int64         = 3
	UserID_3_str string        = "3"
	UserID_3     domain.UserID = 3
	//
	OrderID_i64 int64          = 1
	OrderID_str string         = "1"
	OrderID     domain.OrderID = 1
	//
	ProductID_i64 int64                   = 1
	ProductID_str string                  = "1"
	ProductID     domainProduct.ProductID = 1
	//
	Amount_i64 int64               = 1
	Amount_str string              = "1"
	Amount     domainAmount.Amount = must.Must(domainAmount.New(1))
	//
	AmountPositive_i64 int64                       = 1
	AmountPositive_str string                      = "1"
	AmountPositive     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(1))
	//
	// Use cases
	InValidateToken = must.Must(auth_validate_token.NewInFromValues(AuthToken))
	InGetBalance    = must.Must(get_balance.NewInFromValues(UserID_i64))
)
