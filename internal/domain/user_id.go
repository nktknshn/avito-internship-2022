package domain

import (
	"errors"
	"strconv"
)

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

func NewUserIDFromString(userID string) (UserID, error) {
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return 0, ErrInvalidUserID
	}
	return NewUserID(id)
}
