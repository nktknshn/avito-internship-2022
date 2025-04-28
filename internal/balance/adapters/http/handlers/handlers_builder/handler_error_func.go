package handlers_builder

import (
	"context"
	"net/http"

	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	ergo "github.com/nktknshn/go-ergo-handler"
)

func handlerErrorFunc(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if ergo.IsWrappedError(err) {
		// если ошибка обернута с http статусом, то используем его
		ergo.DefaultHandlerErrorFunc(ctx, w, r, err)
		return
	}

	if domainError.IsDomainError(err) {
		// если ошибка домена, то используем http статус 404
		ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.WrapWithStatusCode(err, http.StatusNotFound))
		return
	}

	// если ошибка не обернута с http статусом, то используем http статус 500
	ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.InternalServerError(err))
}
