package chi

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	commonHttp "github.com/nktknshn/avito-internship-2022/internal/common/http"
)

type ChiRouter struct {
	router *chi.Mux
}

func NewChiRouter() *ChiRouter {
	return &ChiRouter{
		router: chi.NewRouter(),
	}
}

func (r *ChiRouter) GetHandler() http.Handler {
	return r.router
}

func (r *ChiRouter) Attach(route commonHttp.Route) {
	for _, method := range route.GetMethods() {
		r.router.Method(method, route.GetPath(), route.GetHandler())
	}
}

func (r *ChiRouter) Use(middleware ...func(http.Handler) http.Handler) {
	r.router.Use(middleware...)
}

func (r *ChiRouter) GetChi() *chi.Mux {
	return r.router
}

func (r *ChiRouter) Handle(path string, handler http.Handler) {
	r.router.Handle(path, handler)
}
