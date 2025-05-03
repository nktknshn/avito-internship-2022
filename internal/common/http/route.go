package http

import "net/http"

type Route interface {
	GetPath() string
	GetMethods() []string
	GetHandler() http.Handler
}

type RouteAttacher interface {
	Attach(route Route)
}
