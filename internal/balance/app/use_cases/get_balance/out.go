package get_balance

import (
	domainAmount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/amount"
)

type Out struct {
	Available domainAmount.Amount
	Reserved  domainAmount.Amount
}
