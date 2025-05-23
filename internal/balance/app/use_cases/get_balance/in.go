package get_balance

import (
	domain "github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

type In struct {
	userID domain.UserID
}

func NewInFromValues(userID int64) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}
	return In{userID: _userID}, nil
}
