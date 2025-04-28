package report_transactions

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
)

type In struct {
	userID           domain.UserID
	cursor           Cursor
	limit            Limit
	sorting          Sorting
	sortingDirection SortingDirection
	transactionType  TransactionType
}

func NewInFromValues(
	userID int64,
	cursor string,
	limit int,
	sorting string,
	sortingDirection string,
	transactionType string,
) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}

	_cursor, err := NewCursor(cursor)
	if err != nil {
		return In{}, err
	}

	_limit, err := NewLimit(limit)
	if err != nil {
		return In{}, err
	}

	_sorting, err := NewSorting(sorting)
	if err != nil {
		return In{}, err
	}

	_sortingDirection, err := NewSortingDirection(sortingDirection)
	if err != nil {
		return In{}, err
	}

	_transactionType, err := NewTransactionType(transactionType)
	if err != nil {
		return In{}, err
	}

	return In{
		userID:           _userID,
		cursor:           _cursor,
		limit:            _limit,
		sorting:          _sorting,
		sortingDirection: _sortingDirection,
		transactionType:  _transactionType,
	}, nil
}
