package router

import (
	"github.com/gorilla/mux"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers"
)

func CreateMuxRouter(handlers *handlers.Handlers) *mux.Router {
	router := mux.NewRouter()

	router.Handle("/balance/{user_id:[0-9]+}", handlers.GetBalance.GetHandler()).Methods("GET")

	router.Handle("/balance/deposit", handlers.Deposit.GetHandler()).Methods("POST")
	router.Handle("/balance/reserve", handlers.Reserve.GetHandler()).Methods("POST")
	router.Handle("/balance/reserve/cancel", handlers.ReserveCancel.GetHandler()).Methods("POST")
	router.Handle("/balance/reserve/confirm", handlers.ReserveConfirm.GetHandler()).Methods("POST")
	router.Handle("/balance/transfer", handlers.Transfer.GetHandler()).Methods("POST")

	return router
}
