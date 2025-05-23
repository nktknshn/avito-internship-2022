package handlers_builder

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	ergo "github.com/nktknshn/go-ergo-handler"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	domainAccount "github.com/nktknshn/avito-internship-2022/internal/balance/domain/account"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	domainTransaction "github.com/nktknshn/avito-internship-2022/internal/balance/domain/transaction"
)

var (
	errInternal = errors.New("internal server error")
)

func handlerErrorFunc(_ context.Context, w http.ResponseWriter, _ *http.Request, err error) {

	var errorBody any
	var status int

	var internalServerError ergo.InternalServerError
	var errorWithHTTPStatus ergo.ErrorWithHttpStatus

	switch {
	case errors.As(err, &internalServerError):
		status = http.StatusInternalServerError
		errorBody = makeErrorBody(errInternal)
	case errors.As(err, &errorWithHTTPStatus):
		status = errorWithHTTPStatus.HttpStatusCode
		errorBody = makeErrorBody(errorWithHTTPStatus.Err)
	case errors.Is(err, domainAccount.ErrAccountNotFound):
		errorBody = makeErrorBody(err)
		status = http.StatusNotFound
	case errors.Is(err, domainTransaction.ErrTransactionNotFound):
		errorBody = makeErrorBody(err)
		status = http.StatusNotFound
	case errors.Is(err, domainAccount.ErrAccountAlreadyExists):
		errorBody = makeErrorBody(err)
		status = http.StatusConflict
	case errors.Is(err, domainTransaction.ErrTransactionAlreadyPaid):
		errorBody = makeErrorBody(err)
		status = http.StatusConflict
	case errors.Is(err, domainTransaction.ErrTransactionAlreadyReserved):
		errorBody = makeErrorBody(err)
		status = http.StatusConflict
	case errors.Is(err, domainTransaction.ErrTransactionAlreadyCanceled):
		errorBody = makeErrorBody(err)
		status = http.StatusConflict
	case domainError.IsDomainError(err):
		errorBody = makeErrorBody(err)
		status = http.StatusBadRequest
	case useCaseError.IsUseCaseError(err):
		errorBody = makeErrorBody(err)
		status = http.StatusBadRequest
	default:
		errorBody = makeErrorBody(errInternal)
		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	bs, err := json.Marshal(errorBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// TODO: check error
	_, err = w.Write(bs)
	if err != nil {
		//nolint:sloglint // позже придумать, как сделать. Может через контекст?
		slog.Error("error writing response", "error", err)
	}
}
