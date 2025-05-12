package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"

	commonHttp "github.com/nktknshn/avito-internship-2022/internal/common/http"
)

type GorillaRouter struct {
	Router *mux.Router
}

func (r *GorillaRouter) Attach(route commonHttp.Route) {
	handler := route.GetHandler()
	r.Router.Handle(route.GetPath(), handler).Methods(route.GetMethods()...)
}

func (r *GorillaRouter) Use(middleware commonHttp.MiddlewareFunc) {
	r.Router.Use(middleware)
}

func (r *GorillaRouter) GetHandler() http.Handler {
	return r.Router
}

func NewGorillaRouter() *GorillaRouter {
	return &GorillaRouter{
		Router: mux.NewRouter(),
	}
}
