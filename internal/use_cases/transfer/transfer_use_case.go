package transfer

import "github.com/nktknshn/avito-internship-2022/internal/domain"

type In struct {
	From   domain.AccountID
	To     domain.AccountID
	Amount domain.AmountPositive
}
