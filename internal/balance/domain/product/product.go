package product

import domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"

type ProductID int64

func (id ProductID) Value() int64 {
	return int64(id)
}

var (
	ErrInvalidProductID = domainError.New("invalid product id")
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
