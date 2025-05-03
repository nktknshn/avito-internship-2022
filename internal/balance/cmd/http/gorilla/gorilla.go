package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
	commonHttp "github.com/nktknshn/avito-internship-2022/internal/common/http"
)

type gorillaRouter struct {
	Router *mux.Router
}

func (r *gorillaRouter) Attach(route commonHttp.Route) {
	handler := route.GetHandler()
	r.Router.Handle(route.GetPath(), handler).Methods(route.GetMethods()...)
}

func (r *gorillaRouter) Use(middleware commonHttp.MiddlewareFunc) {
	r.Router.Use(middleware)
}

func (r *gorillaRouter) GetHandler() http.Handler {
	return r.Router
}

func NewGorillaRouter() *gorillaRouter {
	return &gorillaRouter{
		Router: mux.NewRouter(),
	}
}
