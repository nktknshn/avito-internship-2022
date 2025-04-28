package handlers_builder

import (
	"context"
	"net/http"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	ergo "github.com/nktknshn/go-ergo-handler"
)

func handlerErrorFunc(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if ergo.IsWrappedError(err) {
		ergo.DefaultHandlerErrorFunc(ctx, w, r, err)
		return
	}
	if domainError.IsDomainError(err) {
		ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.WrapWithStatusCode(err, http.StatusNotFound))
		return
	}
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
