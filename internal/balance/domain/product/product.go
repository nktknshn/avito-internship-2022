package product

import domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"

type ProductID int64

func (id ProductID) Value() int64 {
	return int64(id)
}

var (
	ErrInvalidProductID    = domainError.New("invalid product id")
	ErrInvalidProductTitle = domainError.New("invalid product title")
)

func NewProductID(id int64) (ProductID, error) {
	if id <= 0 {
		return 0, ErrInvalidProductID
	}
	return ProductID(id), nil
}

type ProductTitle string

func (t ProductTitle) Value() string {
	return string(t)
}

func NewProductTitle(title string) (ProductTitle, error) {
	// if len(title) == 0 {
	// 	return "", ErrInvalidProductTitle
	// }
	return ProductTitle(title), nil
}
