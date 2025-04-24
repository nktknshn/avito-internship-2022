package reserve

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
)

type In struct {
	userID    domain.UserID
	productID domainProduct.ProductID
	orderID   domain.OrderID
	amount    domainAmount.AmountPositive
}

func NewInFromValues(userID int64, productID int64, orderID int64, amount int64) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}

	_productID, err := domainProduct.NewProductID(productID)
	if err != nil {
		return In{}, err
	}

	_orderID, err := domain.NewOrderID(orderID)
	if err != nil {
		return In{}, err
	}

	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{
		userID:    _userID,
		productID: _productID,
		orderID:   _orderID,
		amount:    _amount,
	}, nil
}
