package domain

import "errors"

type ProductID int64

func (id ProductID) Value() int64 {
	return int64(id)
}

var (
	ErrInvalidProductID = errors.New("invalid product id")
)

func NewProductID(id int64) (ProductID, error) {
	if id <= 0 {
		return 0, ErrInvalidProductID
	}
	return ProductID(id), nil
}

type ProductName string

type Product struct {
	ID   ProductID
	Name ProductName
}
