package http

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/nktknshn/avito-internship-2022-bench/client_http"
	"github.com/nktknshn/avito-internship-2022-bench/fixtures"
	"github.com/nktknshn/avito-internship-2022-bench/logger"
)

func orderID() int {
	return rand.IntN(100000)
}

type Bench struct {
	client *client_http.ClientWithResponses
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
	_, err := b.client.GetBalanceWithResponse(ctx, userID, client_http.GetBalanceJSONRequestBody{})

	if err != nil {
		logger.GetLogger().Error("getBalance", "error", err)
	}
}

func (b *Bench) deposit(ctx context.Context, userID int, amount int) {
	source := fixtures.RandomDepositSource()
	_, err := b.client.DepositWithResponse(ctx, client_http.DepositJSONRequestBody{
		Amount: &amount,
		UserId: &userID,
		Source: &source,
	})

	if err != nil {
		logger.GetLogger().Error("deposit", "error", err)
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

func (b *Bench) reserve(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	_, err := b.client.ReserveWithResponse(ctx, client_http.ReserveJSONRequestBody{
		Amount:       &amount,
		OrderId:      &orderID,
		ProductId:    &productID,
		ProductTitle: &productTitle,
		UserId:       &userID,
	})

	if err != nil {
		logger.GetLogger().Error("reserve", "error", err)
	}
}

func (b *Bench) reserveConfirm(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	_, err := b.client.ReserveConfirmWithResponse(ctx, client_http.ReserveConfirmJSONRequestBody{
		OrderId:   &orderID,
		ProductId: &productID,
		UserId:    &userID,
		Amount:    &amount,
	})

	if err != nil {
		logger.GetLogger().Error("reserveConfirm", "error", err)
	}
}

func (b *Bench) reserveCancel(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	_, err := b.client.ReserveCancelWithResponse(ctx, client_http.ReserveCancelJSONRequestBody{
		OrderId:   &orderID,
		ProductId: &productID,
		UserId:    &userID,
		Amount:    &amount,
	})

	if err != nil {
		logger.GetLogger().Error("reserveCancel", "error", err)
	}
}

func (b *Bench) transfer(ctx context.Context, fromUserID int, toUserID int, amount int) {
	_, err := b.client.TransferWithResponse(ctx, client_http.TransferJSONRequestBody{
		Amount:     &amount,
		FromUserId: &fromUserID,
		ToUserId:   &toUserID,
	})

	if err != nil {
		logger.GetLogger().Error("transfer", "error", err)
	}
}
