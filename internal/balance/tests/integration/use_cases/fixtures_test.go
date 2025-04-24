package use_cases_test

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

var (
	userID    domain.UserID           = 1
	orderID   domain.OrderID          = 1
	productID domainProduct.ProductID = 1
	amount100                         = fixtures.AmountPositive100
	amount50                          = fixtures.AmountPositive50
)
