package fixtures

import (
	"time"

	"github.com/google/uuid"
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	Amount_i64 int64               = 1
	Amount_str string              = "1"
	Amount     domainAmount.Amount = must.Must(domainAmount.New(1))
	//
	Amount0_i64 int64               = 0
	Amount0_str string              = "0"
	Amount0     domainAmount.Amount = must.Must(domainAmount.New(0))
	//
	Amount100_i64 int64               = 100
	Amount100_str string              = "100"
	Amount100     domainAmount.Amount = must.Must(domainAmount.New(100))
	//
	AmountPositive50_i64 int64                       = 50
	AmountPositive50_str string                      = "50"
	AmountPositive50     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(50))
	//
	AmountPositive_i64 int64                       = 1
	AmountPositive_str string                      = "1"
	AmountPositive     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(1))
	//
	AmountPositive100_i64 int64                       = 100
	AmountPositive100_str string                      = "100"
	AmountPositive100     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(100))
	//
	UUID_1     uuid.UUID = uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	UUID_1_str string    = "123e4567-e89b-12d3-a456-426614174000"
	//
	UUID_2     uuid.UUID = uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")
	UUID_2_str string    = "123e4567-e89b-12d3-a456-426614174001"
	//
	Time_1     time.Time = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	Time_1_str string    = Time_1.Format(time.RFC3339)
)
