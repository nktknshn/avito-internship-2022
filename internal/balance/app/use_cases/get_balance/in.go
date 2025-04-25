package get_balance

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

type In struct {
	UserID domain.UserID
}

func NewInFromValues(userID int64) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}
	return In{UserID: _userID}, nil
}

type Out struct {
	Available int64
	Reserved  int64
}
