package fixtures

import (
	"time"

	"github.com/google/uuid"

	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	Amount_i64 int64               = 1
	Amount_str                     = "1"
	Amount     domainAmount.Amount = must.Must(domainAmount.New(1))
	//
	Amount0_i64 int64
	Amount0_str                     = "0"
	Amount0     domainAmount.Amount = must.Must(domainAmount.New(0))
	//
	Amount50_i64 int64               = 50
	Amount50_str                     = "50"
	Amount50     domainAmount.Amount = must.Must(domainAmount.New(50))
	//
	Amount100_i64 int64               = 100
	Amount100_str                     = "100"
	Amount100     domainAmount.Amount = must.Must(domainAmount.New(100))
	//
	AmountPositive50_i64 int64                       = 50
	AmountPositive50_str                             = "50"
	AmountPositive50     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(50))
	//
	AmountPositive_i64 int64                       = 1
	AmountPositive_str                             = "1"
	AmountPositive     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(1))
	//
	AmountPositive100_i64 int64                       = 100
	AmountPositive100_str                             = "100"
	AmountPositive100     domainAmount.AmountPositive = must.Must(domainAmount.NewPositive(100))
	//
	UUID_1     uuid.UUID = uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	UUID_1_str           = "123e4567-e89b-12d3-a456-426614174000"
	//
	UUID_2     uuid.UUID = uuid.MustParse("123e4567-e89b-12d3-a456-426614174001")
	UUID_2_str           = "123e4567-e89b-12d3-a456-426614174001"
	//
	Time_1     = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	Time_1_str = Time_1.Format(time.RFC3339)
)
