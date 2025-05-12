package handlers_builder

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/require"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
)

func TestHandlerErrorFunc_InternalServerError(t *testing.T) {
	rec := httptest.NewRecorder()
	handlerErrorFunc(t.Context(), rec, nil, ergo.NewInternalServerError(errors.New("test")))

	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.JSONEq(t, `{"error":"internal server error"}`, rec.Body.String())
}

func TestHandlerErrorFunc_WrappedError(t *testing.T) {
	rec := httptest.NewRecorder()
	handlerErrorFunc(t.Context(), rec, nil, ergo.NewError(http.StatusInternalServerError, errors.New("test")))

	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.JSONEq(t, `{"error":"test"}`, rec.Body.String())
}

func TestHandlerErrorFunc_DomainError(t *testing.T) {
	err := domainError.New("domain error")
	rec := httptest.NewRecorder()
	handlerErrorFunc(t.Context(), rec, nil, err)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.JSONEq(t, `{"error":"domain error"}`, rec.Body.String())
}

func TestHandlerErrorFunc_UseCaseError(t *testing.T) {
	err := useCaseError.New("use case error")
	rec := httptest.NewRecorder()
	handlerErrorFunc(t.Context(), rec, nil, err)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.JSONEq(t, `{"error":"use case error"}`, rec.Body.String())
}
