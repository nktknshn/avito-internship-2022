package reserve_cancel

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

type ReserveCancelUseCase struct {
	trm              trm.Manager
	accountRepo      domain.AccountRepository
	transactionsRepo domain.TransactionRepository
}

func NewReserveCancelUseCase(
	trm trm.Manager,
	accountRepo domain.AccountRepository,
	transactionsRepo domain.TransactionRepository,
) *ReserveCancelUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	if transactionsRepo == nil {
		panic("transactionsRepo == nil")
	}

	return &ReserveCancelUseCase{
		trm,
		accountRepo,
		transactionsRepo,
	}
}

// Принимает id пользователя, ИД услуги, ИД заказа, сумму.
type In struct {
	UserID    domain.UserID
	OrderID   domain.OrderID
	ProductID domain.ProductID
	Amount    domain.AmountPositive
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

	_productID, err := domain.NewProductID(productID)
	if err != nil {
		return In{}, err
	}

	_amount, err := domain.NewAmountPositive(amount)
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

func (u *ReserveCancelUseCase) Handle(ctx context.Context, in In) error {
	err := u.trm.Do(ctx, func(ctx context.Context) error {
		acc, err := u.accountRepo.GetByUserID(ctx, in.UserID)
		if err != nil {
			return err
		}

		orderTransactions, err := u.transactionsRepo.GetTransactionSpendByOrderID(ctx, in.UserID, in.OrderID)

		if err != nil {
			return err
		}

		var transaction *domain.TransactionSpend

		for _, transaction = range orderTransactions {
			if transaction.Status == domain.TransactionSpendStatusReserved {
				break
			}
		}

		if transaction == nil {
			return domain.ErrTransactionNotFound
		}

		if transaction.Amount != in.Amount {
			return domain.ErrTransactionAmountMismatch
		}

		err = transaction.Cancel(time.Now())

		if err != nil {
			return err
		}

		_, err = u.transactionsRepo.SaveTransactionSpend(ctx, transaction)

		if err != nil {
			return err
		}

		err = acc.BalanceReserveCancel(in.Amount)

		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, acc)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
