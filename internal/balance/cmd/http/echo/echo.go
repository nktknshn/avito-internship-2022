package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	commonHttp "github.com/nktknshn/avito-internship-2022/internal/common/http"
)

type EchoRouter struct {
	echo *echo.Echo
}

func NewEchoRouter() *EchoRouter {
	return &EchoRouter{}
}

func (e *EchoRouter) GetEcho() *echo.Echo {
	return e.echo
}

func (e *EchoRouter) Use(middleware ...echo.MiddlewareFunc) {
	e.echo.Use(middleware...)
}

func (r *EchoRouter) GetHandler() http.Handler {
	return r.echo
}

func (r *EchoRouter) Attach(route commonHttp.Route) {
	for _, method := range route.GetMethods() {
		r.echo.Add(method, route.GetPath(), echo.WrapHandler(route.GetHandler()))
	}
}
