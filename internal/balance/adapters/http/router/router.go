package router

import (
	"net/http"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	commonHttp "github.com/nktknshn/avito-internship-2022/internal/common/http"
)

type Route struct {
	Path    string
	Methods []string
	Handler http.Handler
}

func (r Route) GetPath() string {
	return r.Path
}

func (r Route) GetMethods() []string {
	return r.Methods
}

func (r Route) GetHandler() http.Handler {
	return r.Handler
}

func Attach(router commonHttp.RouteAttacher, handlers *adaptersHttp.HttpAdapter) {
	router.Attach(Route{
		Path:    "/v1/signin",
		Methods: []string{http.MethodPost},
		Handler: handlers.AuthSignIn.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/balance/{user_id:[0-9]+}",
		Methods: []string{http.MethodGet},
		Handler: handlers.GetBalance.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/balance/deposit",
		Methods: []string{http.MethodPost},
		Handler: handlers.Deposit.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/balance/reserve",
		Methods: []string{http.MethodPost},
		Handler: handlers.Reserve.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/balance/reserve/cancel",
		Methods: []string{http.MethodPost},
		Handler: handlers.ReserveCancel.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/balance/reserve/confirm",
		Methods: []string{http.MethodPost},
		Handler: handlers.ReserveConfirm.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/balance/transfer",
		Methods: []string{http.MethodPost},
		Handler: handlers.Transfer.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/report/transactions/{user_id:[0-9]+}",
		Methods: []string{http.MethodGet},
		Handler: handlers.ReportTransactions.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/report/revenue",
		Methods: []string{http.MethodGet},
		Handler: handlers.ReportRevenue.GetHandler(),
	})

	router.Attach(Route{
		Path:    "/v1/report/revenue/export",
		Methods: []string{http.MethodGet},
		Handler: handlers.ReportRevenueExport.GetHandler(),
	})

}
