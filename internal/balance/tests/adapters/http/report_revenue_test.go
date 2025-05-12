package http_test

import (
	"errors"
	"net/http"

	"github.com/stretchr/testify/mock"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers/handlers_auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue"
	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/report_revenue_export"
	domainAuth "github.com/nktknshn/avito-internship-2022/internal/balance/domain/auth"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/fixtures"
)

func (s *HTTPTestSuite) TestReportRevenue() {
	var validURL = "?year=2021&month=1"

	testCases := []testCase{
		{
			name:       "success",
			auth:       true,
			url:        validURL,
			expectCode: http.StatusOK,
			useCaseReturn: returnSuccess2(report_revenue.Out{
				Records: []report_revenue.OutRecord{
					{
						ProductID:    fixtures.ProductID,
						ProductTitle: fixtures.ProductTitle,
						TotalRevenue: fixtures.Amount100_i64,
					},
					{
						ProductID:    fixtures.ProductID_2,
						ProductTitle: fixtures.ProductTitle_2,
						TotalRevenue: fixtures.Amount100_i64,
					},
				},
			}),
			expectBody: map[string]any{
				"records": []any{
					map[string]any{
						"product_id":    fixtures.ProductID_i64,
						"product_title": fixtures.ProductTitle_str,
						"total_revenue": fixtures.Amount100_i64,
					},
					map[string]any{
						"product_id":    fixtures.ProductID_2_i64,
						"product_title": fixtures.ProductTitle_2_str,
						"total_revenue": fixtures.Amount100_i64,
					},
				},
			},
		},
		{
			name:       "invalid year string",
			auth:       true,
			url:        "?year=invalid_year&month=1",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid int value: invalid_year",
		},
		{
			name:       "invalid month string",
			auth:       true,
			url:        "?year=2021&month=invalid_month",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid int value: invalid_month",
		},
		{
			name:       "invalid year value",
			auth:       true,
			url:        "?year=0&month=1",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid year",
		},
		{
			name:       "invalid month value",
			auth:       true,
			url:        "?year=2021&month=13",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid month",
		},
		{
			name:       "user is not allowed",
			auth:       true,
			url:        validURL,
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
			authRole:   domainAuth.AuthUserRoleNobody,
		},
		{
			name:          "use case error",
			auth:          true,
			url:           validURL,
			expectCode:    http.StatusInternalServerError,
			useCaseReturn: returnError2[report_revenue.Out](errors.New("some error")),
			expectErr:     "internal server error",
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.ReportRevenueUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.ReportRevenue
	}, testCases)
}

func (s *HTTPTestSuite) TestReportRevenueExport() {
	var validURL = "?year=2021&month=1"

	testCases := []testCase{
		{
			name:       "success",
			auth:       true,
			url:        validURL,
			expectCode: http.StatusOK,
			useCaseReturn: returnSuccess2(report_revenue_export.Out{
				URL: "/download/report.csv",
			}),
			expectBody: map[string]any{
				"url": "/download/report.csv",
			},
		},
		{
			name:       "invalid year string",
			auth:       true,
			url:        "?year=invalid_year&month=1",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid int value: invalid_year",
		},
		{
			name:       "invalid month string",
			auth:       true,
			url:        "?year=2021&month=invalid_month",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid int value: invalid_month",
		},
		{
			name:       "invalid year value",
			auth:       true,
			url:        "?year=0&month=1",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid year",
		},
		{
			name:       "invalid month value",
			auth:       true,
			url:        "?year=2021&month=13",
			expectCode: http.StatusBadRequest,
			expectErr:  "invalid month",
		},
		{
			name:       "user is not allowed",
			auth:       true,
			url:        validURL,
			expectCode: http.StatusForbidden,
			expectErr:  handlers_auth.ErrUserNotAllowed.Error(),
			authRole:   domainAuth.AuthUserRoleNobody,
		},
		{
			name:          "use case error",
			auth:          true,
			url:           validURL,
			expectCode:    http.StatusInternalServerError,
			useCaseReturn: returnError2[report_revenue_export.Out](errors.New("some error")),
			expectErr:     "internal server error",
		},
	}

	s.runTestCases(func() *mock.Mock {
		return &s.app.ReportRevenueExportUseCaseMock.Mock
	}, func() adaptersHttp.Handler {
		return s.httpAdapter.ReportRevenueExport
	}, testCases)
}
