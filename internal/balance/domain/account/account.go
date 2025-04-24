package account

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain"
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

var (
	ErrSameAccount             = domainError.New("same account")
	ErrFailedToRollbackBalance = domainError.New("failed to rollback balance")
)

type AccountID int64

func (a AccountID) Value() int64 {
	return int64(a)
}

var (
	ErrInvalidAccountID = domainError.New("invalid account id")
)

func NewAccountID(id int64) (AccountID, error) {
	if id < 0 {
		return 0, ErrInvalidAccountID
	}
	return AccountID(id), nil
}

type Account struct {
	ID      AccountID
	UserID  domain.UserID
	Balance AccountBalance
}

func NewAccount(userID domain.UserID) (*Account, error) {
	return &Account{
		UserID:  userID,
		Balance: NewAccountBalanceEmpty(),
	}, nil
}

func NewAccountFromValues(id int64, userID int64, balanceAvailable int64, balanceReserved int64) (*Account, error) {
	_id, err := NewAccountID(id)
	if err != nil {
		return nil, err
	}
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return nil, err
	}
	_balanceAvailable, err := amount.New(balanceAvailable)
	if err != nil {
		return nil, err
	}
	_balanceReserved, err := amount.New(balanceReserved)
	if err != nil {
		return nil, err
	}
	accountBalance, err := NewAccountBalance(_balanceAvailable, _balanceReserved)
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:      _id,
		UserID:  _userID,
		Balance: accountBalance,
	}, nil
}

func (ac *Account) SetBalance(ab AccountBalance) error {
	ac.Balance = ab
	return nil
}

func (ac *Account) BalanceDeposit(amount amount.AmountPositive) error {
	newBalance, err := ac.Balance.Deposit(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserve(a amount.AmountPositive) error {
	newBalance, err := ac.Balance.Reserve(a)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserveCancel(a amount.AmountPositive) error {
	newBalance, err := ac.Balance.ReserveCancel(a)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceReserveConfirm(a amount.AmountPositive) error {
	newBalance, err := ac.Balance.ReserveConfirm(a)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) BalanceWithdraw(amount amount.AmountPositive) error {
	newBalance, err := ac.Balance.Withdraw(amount)
	if err != nil {
		return err
	}
	err = ac.SetBalance(newBalance)
	if err != nil {
		return err
	}
	return nil
}

func (ac *Account) Transfer(to *Account, amount amount.AmountPositive) error {
	if ac.ID == to.ID {
		return ErrSameAccount
	}

	originalBalance := ac.Balance

	err := ac.BalanceWithdraw(amount)

	if err != nil {
		return err
	}

	err = to.BalanceDeposit(amount)

	if err != nil {
		err = ac.SetBalance(originalBalance)
		if err != nil {
			return ErrFailedToRollbackBalance
		}
		return err
	}

	return nil
}
