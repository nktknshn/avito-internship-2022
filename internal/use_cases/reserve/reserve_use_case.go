package reserve

import (
	"context"
	"time"

	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/nktknshn/avito-internship-2022/internal/domain"
)

type ReserveUseCase struct {
	trm             trm.Manager
	accountRepo     domain.AccountRepository
	transactionRepo domain.TransactionRepository
}

func NewReserveUseCase(
	trm trm.Manager,
	accountRepo domain.AccountRepository,
	transactionRepo domain.TransactionRepository,
) *ReserveUseCase {

	if trm == nil {
		panic("trm == nil")
	}

	if accountRepo == nil {
		panic("accountRepo == nil")
	}

	if transactionRepo == nil {
		panic("transactionRepo == nil")
	}

	return &ReserveUseCase{
		trm,
		accountRepo,
		transactionRepo,
	}
}

type In struct {
	UserID    domain.UserID
	ProductID domain.ProductID
	OrderID   domain.OrderID
	Amount    domain.AmountPositive
}

func NewInFromValues(userID int64, productID int64, orderID int64, amount int64) (In, error) {
	_userID, err := domain.NewUserID(userID)
	if err != nil {
		return In{}, err
	}

	_productID, err := domain.NewProductID(productID)
	if err != nil {
		return In{}, err
	}

	_orderID, err := domain.NewOrderID(orderID)
	if err != nil {
		return In{}, err
	}

	_amount, err := domain.NewAmountPositive(amount)
	if err != nil {
		return In{}, err
	}

	return In{
		UserID:    _userID,
		ProductID: _productID,
		OrderID:   _orderID,
		Amount:    _amount,
	}, nil
}

func (u *ReserveUseCase) Handle(ctx context.Context, in In) error {

	// а если canceled, то OrderID новый?
	// если есть резерв с OrderID и статус не canceled, то ошибка

	err := u.trm.Do(ctx, func(ctx context.Context) error {

		acc, err := u.accountRepo.GetByUserID(ctx, in.UserID)

		if err != nil {
			return err
		}

		orderTransactions, err := u.transactionRepo.GetTransactionSpendByOrderID(ctx, in.UserID, in.OrderID)

		if err != nil {
			return err
		}

		for _, transaction := range orderTransactions {
			// если существует транзакция с таким OrderID и статус не canceled, то ошибка
			if transaction.Status != domain.TransactionSpendStatusCanceled {
				return domain.ErrTransactionAlreadyExists
			}
		}

		err = acc.BalanceReserve(in.Amount)

		if err != nil {
			return err
		}

		_, err = u.accountRepo.Save(ctx, acc)

		if err != nil {
			return err
		}

		transaction, err := domain.NewTransactionSpendReserved(
			acc.ID,
			in.UserID,
			in.OrderID,
			in.ProductID,
			in.Amount,
			time.Now(),
		)

		if err != nil {
			return err
		}

		_, err = u.transactionRepo.SaveTransactionSpend(ctx, transaction)

		if err != nil {
			return err
		}

		return nil
	})

	return err
}
