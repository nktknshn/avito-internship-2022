package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

var (
	AuthToken              string                  = "token"
	AuthUserID_i64         int64                   = 1
	AuthUserID_str         string                  = "1"
	AuthUserID             domainAuth.AuthUserID   = 1
	AuthUserID_i64_invalid int64                   = -1
	AuthUserID_str_invalid string                  = "-1"
	AuthUserRole_str       string                  = "admin"
	AuthUserRole           domainAuth.AuthUserRole = domainAuth.AuthUserRoleAdmin
	//
	UsernameAdmin_str     string                          = "admin"
	UsernameAdmin         domainAuth.AuthUserUsername     = domainAuth.AuthUserUsername(UsernameAdmin_str)
	PasswordAdmin_str     string                          = "password123"
	PasswordAdmin         domainAuth.AuthUserPassword     = domainAuth.AuthUserPassword(PasswordAdmin_str)
	PasswordHashAdmin_str string                          = "JGFyZ29uMmlkJHY9MTkkbT02NTUzNix0PTEscD00JFRHNVZkMFJVTkhWUFFqazJNR3B3YWckWHRUSG5xVVFlUmpjWFRWZ0NUSGZQeEFPbm9BYThaREpkOFIxdkhUTDVEcw=="
	PasswordHashAdmin     domainAuth.AuthUserPasswordHash = domainAuth.AuthUserPasswordHash(PasswordHashAdmin_str)
	//
	AuthUser domainAuth.AuthUser = domainAuth.AuthUser{
		ID:           AuthUserID,
		Username:     UsernameAdmin,
		PasswordHash: PasswordHashAdmin,
		Role:         AuthUserRole,
	}
	//
	AccountID_i64         int64                   = 1
	AccountID_str         string                  = "1"
	AccountID             domainAccount.AccountID = 1
	AccountID_2_i64       int64                   = 2
	AccountID_2_str       string                  = "2"
	AccountID_2           domainAccount.AccountID = 2
	AccountID_i64_invalid int64                   = -1
	AccountID_str_invalid string                  = "-1"
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
	UserID_i64_invalid int64  = -1
	UserID_str_invalid string = "-1"
	//
	OrderID_i64 int64          = 1
	OrderID_str string         = "1"
	OrderID     domain.OrderID = 1
	//
	ProductID_i64   int64                   = 1
	ProductID_str   string                  = "1"
	ProductID       domainProduct.ProductID = 1
	ProductID_2_i64 int64                   = 2
	ProductID_2_str string                  = "2"
	ProductID_2     domainProduct.ProductID = 2
	//
	ProductTitle_str string                     = "Product Title"
	ProductTitle     domainProduct.ProductTitle = domainProduct.ProductTitle(ProductTitle_str)
	//
	ProductTitle_2_str string                     = "Product Title 2"
	ProductTitle_2     domainProduct.ProductTitle = domainProduct.ProductTitle(ProductTitle_2_str)
	//
	DepositSource_str string                                     = "credit_card"
	DepositSource     domainTransaction.TransactionDepositSource = domainTransaction.TransactionDepositSource(DepositSource_str)
)
