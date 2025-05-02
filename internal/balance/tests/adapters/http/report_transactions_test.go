package http_test

import (
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_transactions"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
	"github.com/stretchr/testify/mock"
)

func (s *HttpTestSuite) TestReportTransactions() {

	var validRouteParams = map[string]string{
		"user_id": fixtures.UserID_str,
	}

	var validUrl = "?limit=10&sorting=updated_at&sorting_direction=desc"

	var transactionSpend = report_transactions.OutTransactionSpend{
		ID:           domainTransaction.TransactionSpendID(fixtures.UUID_1),
		AccountID:    fixtures.AccountID,
		OrderID:      fixtures.OrderID,
		ProductID:    fixtures.ProductID,
		ProductTitle: fixtures.ProductTitle,
		Amount:       fixtures.AmountPositive100,
		Status:       domainTransaction.TransactionSpendStatusConfirmed,
		CreatedAt:    fixtures.Time_1,
		UpdatedAt:    fixtures.Time_1,
	}

	testCases := []testCase{
		{
			name:        "success",
			url:         validUrl,
			routeParams: validRouteParams,
			useCaseReturn: returnSuccess2(report_transactions.Out{
				Cursor:  "cursor",
				HasMore: true,
				Transactions: []report_transactions.OutTransaction{
					&transactionSpend,
				},
			}),
			expectCode: http.StatusOK,
			expectBody: map[string]any{
				"transactions": []any{
					map[string]any{
						"id":            fixtures.UUID_1_str,
						"account_id":    fixtures.AccountID_i64,
						"order_id":      fixtures.OrderID_i64,
						"product_id":    fixtures.ProductID_i64,
						"product_title": fixtures.ProductTitle_str,
						"amount":        fixtures.Amount100_i64,
						"status":        domainTransaction.TransactionSpendStatusConfirmed.Value(),
						"created_at":    fixtures.Time_1_str,
						"updated_at":    fixtures.Time_1_str,
					},
				},
				"cursor":   "cursor",
				"has_more": true,
			},
			auth: true,
		},
		{
			name:          "not found",
			auth:          true,
			url:           validUrl,
			routeParams:   validRouteParams,
			useCaseReturn: returnError2[report_transactions.Out](domainAccount.ErrAccountNotFound),
			expectCode:    http.StatusNotFound,
			expectErr:     domainAccount.ErrAccountNotFound.Error(),
		},
		{
			name:       "user is not allowed",
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
			auth:       true,
			authRole:   domainAuth.AuthUserRoleNobody,
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.ReportTransactionsUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.ReportTransactions
	}, testCases)
}
