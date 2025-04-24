package reserve_confirm

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
)

type In struct {
	UserID    domain.UserID
	OrderID   domain.OrderID
	ProductID domainProduct.ProductID
	Amount    domainAmount.AmountPositive
}

func NewInFromValues(userID int64, orderID int64, productID int64, amount int64) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}

	_orderID, err := domain.NewOrderID(orderID)
	if err != nil {
		return In{}, err
	}

	_productID, err := domainProduct.NewProductID(productID)
	if err != nil {
		return In{}, err
	}

	_amount, err := domainAmount.NewPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{
		UserID:    _userID,
		OrderID:   _orderID,
		ProductID: _productID,
		Amount:    _amount,
	}, nil
}
