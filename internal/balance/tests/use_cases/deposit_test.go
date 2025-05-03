package use_cases_test

import (
	"context"
	"errors"
	"fmt"

	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *SuiteTest) TestDeposit_Errors() {

	s.mockedAccountsRepo.On("GetByUserID", mock.Anything, mock.Anything).Return(nil, domainAccount.ErrAccountNotFound)
	s.mockedAccountsRepo.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("save error"))

	err := s.mockedDeposit.Handle(context.Background(), fixtures.InDeposit100)

	s.Require().Error(err)
	fmt.Println(err)

}
