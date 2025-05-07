package http

import (
	"context"
	"io"
	"math/rand/v2"
	"net/http"
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

func readBody(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return []byte("nil"), nil
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

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
	_, resp, err := req.Execute()
	if err != nil {
		respBody, _ := readBody(resp)
		logger.GetLogger().Error("getBalance", "error", err, "body", respBody, "userID", userID)
	}
}

func (b *Bench) deposit(ctx context.Context, userID int, amount int) error {
	req := b.c2.DepositAPI.Deposit(ctx)
	amount32 := int32(amount)
	userID32 := int32(userID)
	source := fixtures.RandomDepositSource()
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersDepositRequestBody{
		Amount: &amount32,
		Source: &source,
		UserId: &userID32,
	})
	_, resp, err := req.Execute()
	if err != nil {
		respBody, _ := readBody(resp)
		logger.GetLogger().Error("deposit", "error", err, "body", respBody, "userID", userID, "amount", amount, "source", source)
		return err
	}
	return nil
}

func (b *Bench) reserve(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) error {
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
	_, resp, err := req.Execute()
	if err != nil {
		respBody, _ := readBody(resp)
		logger.GetLogger().Error("reserve", "error", err, "body", respBody, "userID", userID, "orderID", orderID, "productID", productID, "productTitle", productTitle, "amount", amount)
		return err
	}
	return nil
}

func (b *Bench) reserveConfirm(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) error {
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

	_, resp, err := req.Execute()

	if err != nil {
		respBody, _ := readBody(resp)
		logger.GetLogger().Error("reserveConfirm", "error", err, "body", respBody, "userID", userID, "orderID", orderID, "productID", productID, "productTitle", productTitle, "amount", amount)
		return err
	}
	return nil
}

func (b *Bench) reserveCancel(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) error {
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
	_, resp, err := req.Execute()

	if err != nil {
		respBody, _ := readBody(resp)
		logger.GetLogger().Error("reserveCancel", "error", err, "body", respBody, "userID", userID, "orderID", orderID, "productID", productID, "productTitle", productTitle, "amount", amount)
		return err
	}
	return nil
}

func (b *Bench) transfer(ctx context.Context, fromUserID int, toUserID int, amount int) error {
	req := b.c2.TransferAPI.Transfer(ctx)
	amount32 := int32(amount)
	fromUserID32 := int32(fromUserID)
	toUserID32 := int32(toUserID)
	req = req.Payload(openapi.InternalBalanceAdaptersHttpHandlersTransferRequestBody{
		Amount:     &amount32,
		FromUserId: &fromUserID32,
		ToUserId:   &toUserID32,
	})
	_, resp, err := req.Execute()
	if err != nil {
		respBody, _ := readBody(resp)
		logger.GetLogger().Error("transfer", "error", err, "body", respBody, "fromUserID", fromUserID, "toUserID", toUserID, "amount", amount)
		return err
	}
	return nil
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

func (b *Bench) scenarioReserveSuccess(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	if err := b.deposit(ctx, userID, amount+10); err != nil {
		return
	}
	if err := b.reserve(ctx, userID, orderID, productID, productTitle, amount); err != nil {
		return
	}
	time.Sleep(b.getDelay())
	if err := b.reserveConfirm(ctx, userID, orderID, productID, productTitle, amount); err != nil {
		return
	}
}

func (b *Bench) scenarioReserveCancel(ctx context.Context, userID int, orderID int, productID int, productTitle string, amount int) {
	if err := b.deposit(ctx, userID, amount+10); err != nil {
		return
	}
	if err := b.reserve(ctx, userID, orderID, productID, productTitle, amount); err != nil {
		return
	}
	time.Sleep(b.getDelay())
	if err := b.reserveCancel(ctx, userID, orderID, productID, productTitle, amount); err != nil {
		return
	}
}
