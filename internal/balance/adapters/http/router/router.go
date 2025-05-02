package router

import (
	"net/http"

	"github.com/gorilla/mux"
	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
)

type Route struct {
	Path    string
	Methods []string
	Handler http.Handler
}

func AttachHandlers(router interface{ Attach(route Route) }, handlers *adaptersHttp.HttpAdapter) {
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

}

func NewMuxRouter(handlers *adaptersHttp.HttpAdapter) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/v1/signin", handlers.AuthSignIn.GetHandler()).
		Methods(http.MethodPost)

	router.Handle("/v1/balance/{user_id:[0-9]+}", handlers.GetBalance.GetHandler()).
		Methods(http.MethodGet)

	router.Handle("/v1/balance/deposit", handlers.Deposit.GetHandler()).
		Methods(http.MethodPost)

	router.Handle("/v1/balance/reserve", handlers.Reserve.GetHandler()).
		Methods(http.MethodPost)

	router.Handle("/v1/balance/reserve/cancel", handlers.ReserveCancel.GetHandler()).
		Methods(http.MethodPost)

	router.Handle("/v1/balance/reserve/confirm", handlers.ReserveConfirm.GetHandler()).
		Methods(http.MethodPost)

	router.Handle("/v1/balance/transfer", handlers.Transfer.GetHandler()).
		Methods(http.MethodPost)

	router.Handle("/v1/report/transactions/{user_id:[0-9]+}", handlers.ReportTransactions.GetHandler()).
		Methods(http.MethodGet)

	router.Handle("/v1/report/revenue", handlers.ReportRevenue.GetHandler()).
		Methods(http.MethodGet)

	return router
}
