package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	balanceRouter "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
)

type gorillaRouter struct {
	*mux.Router
	// Middlewares []MiddlewareFunc
}

func (r *gorillaRouter) Attach(route balanceRouter.Route) {
	handler := route.Handler

	r.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	})

	r.Router.Handle(route.Path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(route.Methods, ", "))
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusNoContent)
	})).Methods(http.MethodOptions)

	r.Handle(route.Path, handler).Methods(route.Methods...)
}

func (r *gorillaRouter) Use(middleware MiddlewareFunc) {
	r.Router.Use(middleware)
	// r.Middlewares = append(r.Middlewares, middleware)
}

func NewGorillaRouter() *gorillaRouter {
	return &gorillaRouter{
		Router: mux.NewRouter(),
		// Middlewares: []MiddlewareFunc{},
	}
}
