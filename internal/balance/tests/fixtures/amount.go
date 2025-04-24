package fixtures

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
	"github.com/nktknshn/avito-internship-2022/internal/common/helpers/must"
)

var (
	Amount0           = must.Must(amount.New(0))
	Amount100         = must.Must(amount.New(100))
	AmountPositive100 = must.Must(amount.NewPositive(100))
	AmountPositive50  = must.Must(amount.NewPositive(50))
)
