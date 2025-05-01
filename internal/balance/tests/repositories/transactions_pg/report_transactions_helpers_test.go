package transactions_pg_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainProduct "github.com/nktknshn/avito-internship-2022/internal/balance/domain/product"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

type transactionWrapper struct {
	trSpend    *domainTransaction.TransactionSpend
	trDeposit  *domainTransaction.TransactionDeposit
	trTransfer *domainTransaction.TransactionTransfer
}

func (tw *transactionWrapper) setProduct(p *struct {
	productID   int64
	productName string
}) *transactionWrapper {
	if tw.isSpend() {
		tw.trSpend.ProductID = domainProduct.ProductID(p.productID)
		tw.trSpend.ProductTitle = domainProduct.ProductTitle(p.productName)
	}
	return tw
}

func (tw *transactionWrapper) setStatus(status domainTransaction.TransactionSpendStatus) *transactionWrapper {
	if tw.isSpend() {
		tw.trSpend.Status = status
	}
	return tw
}

func (tw *transactionWrapper) isSpend() bool {
	return tw.trSpend != nil
}

func (tw *transactionWrapper) isDeposit() bool {
	return tw.trDeposit != nil
}

func (tw *transactionWrapper) isTransfer() bool {
	return tw.trTransfer != nil
}

func (tw *transactionWrapper) getID() uuid.UUID {
	if tw.trSpend != nil {
		return tw.trSpend.ID.Value()
	}

	if tw.trDeposit != nil {
		return tw.trDeposit.ID.Value()
	}

	if tw.trTransfer != nil {
		return tw.trTransfer.ID.Value()
	}

	panic("unknown transaction type")
}

func (tw *transactionWrapper) getAmount() int64 {
	if tw.trSpend != nil {
		return tw.trSpend.Amount.Value()
	}

	if tw.trDeposit != nil {
		return tw.trDeposit.Amount.Value()
	}

	if tw.trTransfer != nil {
		return tw.trTransfer.Amount.Value()
	}

	return 0
}

func (tw *transactionWrapper) getUpdatedAt() time.Time {
	if tw.trSpend != nil {
		return tw.trSpend.UpdatedAt
	}

	if tw.trDeposit != nil {
		return tw.trDeposit.UpdatedAt
	}

	if tw.trTransfer != nil {
		return tw.trTransfer.UpdatedAt
	}

	return time.Time{}
}

func tSpend(acc *domainAccount.Account, amount int64, updatedAt time.Time) *transactionWrapper {

	trSpend, err := domainTransaction.NewTransactionSpendFromValues(
		uuid.Nil,
		acc.ID.Value(),
		fixtures.UserID_i64,
		fixtures.OrderID_i64,
		fixtures.ProductID_i64,
		fixtures.ProductTitle_str,
		amount,
		domainTransaction.TransactionSpendStatusConfirmed.Value(),
		updatedAt,
		updatedAt,
	)

	if err != nil {
		panic(err)
	}

	return &transactionWrapper{
		trSpend: trSpend,
	}
}

func tDeposit(acc *domainAccount.Account, amount int64, updatedAt time.Time) *transactionWrapper {

	trDeposit, err := domainTransaction.NewTransactionDepositFromValues(
		uuid.Nil,
		acc.ID.Value(),
		fixtures.UserID_i64,
		fixtures.DepositSource_str,
		domainTransaction.TransactionDepositStatusConfirmed.Value(),
		amount,
		updatedAt,
		updatedAt,
	)

	if err != nil {
		panic(err)
	}

	return &transactionWrapper{
		trDeposit: trDeposit,
	}
}

func tTransfer(accFrom *domainAccount.Account, accTo *domainAccount.Account, amount int64, updatedAt time.Time) *transactionWrapper {

	trTransfer, err := domainTransaction.NewTransactionTransferFromValues(
		uuid.Nil,
		accFrom.ID.Value(),
		accTo.ID.Value(),
		amount,
		domainTransaction.TransactionTransferStatusConfirmed.Value(),
		updatedAt,
		updatedAt,
	)

	if err != nil {
		panic(err)
	}

	return &transactionWrapper{
		trTransfer: trTransfer,
	}
}

func (s *Suite) saveTransactions(trs []*transactionWrapper) ([]*transactionWrapper, error) {

	trs_saved := make([]*transactionWrapper, len(trs))

	for i, tr := range trs {
		switch {
		case tr.isSpend():
			tr_saved, err := s.transactionsRepo.SaveTransactionSpend(s.Context(), tr.trSpend)
			s.Require().NoError(err)
			trs_saved[i] = &transactionWrapper{
				trSpend: tr_saved,
			}
		case tr.isDeposit():
			tr_saved, err := s.transactionsRepo.SaveTransactionDeposit(s.Context(), tr.trDeposit)
			s.Require().NoError(err)
			trs_saved[i] = &transactionWrapper{
				trDeposit: tr_saved,
			}
		case tr.isTransfer():
			tr_saved, err := s.transactionsRepo.SaveTransactionTransfer(s.Context(), tr.trTransfer)
			s.Require().NoError(err)
			trs_saved[i] = &transactionWrapper{
				trTransfer: tr_saved,
			}
		default:
			return nil, fmt.Errorf("unknown transaction type: %T", tr)
		}
	}
	return trs_saved, nil
}

func transactionsEqual(a []*transactionWrapper, b []report_transactions.Transaction) bool {
	if len(a) != len(b) {
		return false
	}

	for i, tr := range a {
		switch tr_b := b[i].(type) {
		case *domainTransaction.TransactionSpend:
			if tr.isSpend() && tr_b.ID.Value() == tr.getID() {
				continue
			}
			return false
		case *domainTransaction.TransactionDeposit:
			if tr.isDeposit() && tr_b.ID.Value() == tr.getID() {
				continue
			}
			return false
		case *domainTransaction.TransactionTransfer:
			if tr.isTransfer() && tr_b.ID.Value() == tr.getID() {
				continue
			}
			return false
		default:
			panic(fmt.Sprintf("unknown transaction type: %T", tr_b))
		}
	}

	return true
}

func rInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func rTime(min, max time.Time) time.Time {
	return min.Add(time.Duration(rand.Int63n(max.Unix()-min.Unix()+1)) * time.Second)
}

func sortAmountAsc(trs []*transactionWrapper) {
	slices.SortFunc(trs, func(a, b *transactionWrapper) int {
		return bytes.Compare(
			[]byte(a.getID().String()),
			[]byte(b.getID().String()),
		)
	})
	slices.SortStableFunc(trs, func(a, b *transactionWrapper) int {
		return int(a.getAmount() - b.getAmount())
	})
}

func sortAmountDesc(trs []*transactionWrapper) {
	sortAmountAsc(trs)
	slices.Reverse(trs)
}

func sortUpdatedAtAsc(trs []*transactionWrapper) {
	slices.SortFunc(trs, func(a, b *transactionWrapper) int {
		return bytes.Compare(
			[]byte(a.getID().String()),
			[]byte(b.getID().String()),
		)
	})
	slices.SortStableFunc(trs, func(a, b *transactionWrapper) int {
		return int(a.getUpdatedAt().Sub(b.getUpdatedAt()))
	})
}

func sortUpdatedAtDesc(trs []*transactionWrapper) {
	sortUpdatedAtAsc(trs)
	slices.Reverse(trs)
}
