package domain

import "errors"

type UserID int64

func (u UserID) Value() int64 {
	return int64(u)
}

var ErrInvalidUserID = errors.New("invalid UserID")

func NewUserID(userID int64) (UserID, error) {
	if userID <= 0 {
		return 0, ErrInvalidUserID
	}
	return UserID(userID), nil
}
