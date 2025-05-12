package server

import (
	"context"
	"net/http"
)

type MiddlewareFunc = func(next http.Handler) http.Handler

type ErrorHandler = func(ctx context.Context, w http.ResponseWriter, r *http.Request, err any)

func NewMiddlewareRecover(errorHandler ErrorHandler) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				//nolint:errorlint // okay to check this
				if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
					errorHandler(r.Context(), w, r, rvr)
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
