package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/get_balance"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
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
	ProductTitle_str string                     = "Product Title"
	ProductTitle     domainProduct.ProductTitle = domainProduct.ProductTitle(ProductTitle_str)
	//
	DepositSource_str string                                     = "credit_card"
	DepositSource     domainTransaction.TransactionDepositSource = domainTransaction.TransactionDepositSource(DepositSource_str)
	//
	// Use cases
	InValidateToken = must.Must(auth_validate_token.NewInFromValues(AuthToken))
	InGetBalance    = must.Must(get_balance.NewInFromValues(UserID_i64))
)
