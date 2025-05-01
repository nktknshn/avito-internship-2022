package handlers_builder

import (
	"context"
	"net/http"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	ergo "github.com/nktknshn/go-ergo-handler"
)

func handlerErrorFunc(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	if ergo.IsWrappedError(err) {
		// если ошибка обернута в хендлере с http статусом, то используем ее
		ergo.DefaultHandlerErrorFunc(ctx, w, r, err)
		return
	}

	if domainError.IsDomainError(err) {
		// если ошибка домена, то используем http статус 400
		ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.WrapWithStatusCode(err, http.StatusBadRequest))
		return
	}

	if useCaseError.IsUseCaseError(err) {
		// если ошибка юзкейса, то используем http статус 400
		ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.WrapWithStatusCode(err, http.StatusBadRequest))
		return
	}

	// если ошибка не обернута с http статусом, то используем http статус 500
	ergo.DefaultHandlerErrorFunc(ctx, w, r, ergo.InternalServerError(err))
}
