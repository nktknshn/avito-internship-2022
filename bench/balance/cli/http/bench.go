package http

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"

	openapi "github.com/nktknshn/avito-internship-2022-bench/api"
	"github.com/nktknshn/avito-internship-2022-bench/fixtures"
	"github.com/nktknshn/avito-internship-2022-bench/logger"
)

func orderID() int {
	return rand.IntN(100000)
}

type Bench struct {
	c2 *openapi.APIClient
}

var delay = 300 * time.Millisecond

func (b *Bench) Bench(ctx context.Context) {
	workresCount := 50

	wg := sync.WaitGroup{}

	for range workresCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.worker()
		}()
	}

	wg.Wait()
}

func (b *Bench) randomScenario(ctx context.Context, userID int) {
	orderID := orderID()
	product := fixtures.RandomProduct()
	amount := rand.IntN(1000)

	switch rand.IntN(5) {
	case 0:
		b.scenarioReserveSuccess(ctx, userID, orderID, product.ID, product.Title, amount)
	case 1:
		b.scenarioReserveCancel(ctx, userID, orderID, product.ID, product.Title, amount)
	case 2:
		b.deposit(ctx, userID, amount)
	case 3:
		b.getBalance(ctx, userID)
	case 4:
		b.transfer(ctx, userID, rand.IntN(1000), amount)
	}
}

func (b *Bench) getDelay() time.Duration {
	return time.Duration(rand.IntN(100)) * time.Millisecond
}

func (b *Bench) worker() {
	for {
		userID := rand.IntN(1000)
		b.randomScenario(context.Background(), userID)
		time.Sleep(b.getDelay())
	}
}

func (b *Bench) getBalance(ctx context.Context, userID int) {
	req := b.c2.BalanceAPI.GetBalance(ctx, int32(userID))
	_, _, err := req.Execute()
	if err != nil {
		logger.GetLogger().Error("getBalance", "error", err)
	}
}

func (b *Bench) deposit(ctx context.Context, userID int, amount int) {
	req := b.c2.DepositAPI.Deposit(ctx)
	amount32 := int32(amount)
	userID32 := int32(userID)
	source := fixtures.RandomDepositSource()
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersDepositRequestBody{
		Amount: &amount32,
		Source: &source,
		UserId: &userID32,
	})
	_, _, err := req.Execute()
	if err != nil {
		logger.GetLogger().Error("deposit", "error", err)
	}
}

func (b *Bench) reserve(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	req := b.c2.ReserveAPI.Reserve(ctx)
	amount32 := int32(amount)
	orderID32 := int32(orderID)
	productID32 := int32(productID)
	userID32 := int32(userID)
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersReserveRequestBody{
		Amount:       &amount32,
		OrderId:      &orderID32,
		ProductId:    &productID32,
		ProductTitle: &productTitle,
		UserId:       &userID32,
	})
	_, _, err := req.Execute()
	if err != nil {
		logger.GetLogger().Error("reserve", "error", err)
	}
}

func (b *Bench) reserveConfirm(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	req := b.c2.ReserveConfirmAPI.ReserveConfirm(ctx)
	amount32 := int32(amount)
	orderID32 := int32(orderID)
	productID32 := int32(productID)
	userID32 := int32(userID)
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersReserveConfirmRequestBody{
		Amount:    &amount32,
		OrderId:   &orderID32,
		ProductId: &productID32,
		UserId:    &userID32,
	})
	_, _, err := req.Execute()
	if err != nil {
		logger.GetLogger().Error("reserveConfirm", "error", err)
	}
}

func (b *Bench) reserveCancel(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	req := b.c2.ReserveCancelAPI.ReserveCancel(ctx)
	amount32 := int32(amount)
	orderID32 := int32(orderID)
	productID32 := int32(productID)
	userID32 := int32(userID)
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersReserveCancelRequestBody{
		Amount:    &amount32,
		OrderId:   &orderID32,
		ProductId: &productID32,
		UserId:    &userID32,
	})
	_, _, err := req.Execute()
	if err != nil {
		logger.GetLogger().Error("reserveCancel", "error", err)
	}
}

func (b *Bench) transfer(ctx context.Context, fromUserID int, toUserID int, amount int) {
	req := b.c2.TransferAPI.Transfer(ctx)
	amount32 := int32(amount)
	fromUserID32 := int32(fromUserID)
	toUserID32 := int32(toUserID)
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersTransferRequestBody{
		Amount:     &amount32,
		FromUserId: &fromUserID32,
		ToUserId:   &toUserID32,
	})
	_, _, err := req.Execute()
	if err != nil {
		logger.GetLogger().Error("transfer", "error", err)
	}
}

func (b *Bench) scenarioReserveSuccess(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	b.deposit(ctx, userID, amount+10)
	b.reserve(ctx, userID, orderID, productID, productTitle, amount)
	time.Sleep(b.getDelay())
	b.reserveConfirm(ctx, userID, orderID, productID, productTitle, amount)
}

func (b *Bench) scenarioReserveCancel(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	b.deposit(ctx, userID, amount+10)
	b.reserve(ctx, userID, orderID, productID, productTitle, amount)
	time.Sleep(b.getDelay())
	b.reserveCancel(ctx, userID, orderID, productID, productTitle, amount)
}
