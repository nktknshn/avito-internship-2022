package handlers_builder

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	useCaseError "github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/errors"
	domainError "github.com/nktknshn/avito-internship-2022/internal/balance/domain/errors"
	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/require"
)

func TestHandlerErrorFunc_WrappedError(t *testing.T) {
	rec := httptest.NewRecorder()
	handlerErrorFunc(context.Background(), rec, nil, ergo.NewError(http.StatusInternalServerError, errors.New("test")))

	require.Equal(t, http.StatusInternalServerError, rec.Code)
	require.Equal(t, `{"error":"test"}`, rec.Body.String())
}

func TestHandlerErrorFunc_DomainError(t *testing.T) {
	err := domainError.New("domain error")
	rec := httptest.NewRecorder()
	handlerErrorFunc(context.Background(), rec, nil, err)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Equal(t, `{"error":"domain error"}`, rec.Body.String())
}

func TestHandlerErrorFunc_UseCaseError(t *testing.T) {
	err := useCaseError.New("use case error")
	rec := httptest.NewRecorder()
	handlerErrorFunc(context.Background(), rec, nil, err)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Equal(t, `{"error":"use case error"}`, rec.Body.String())
}
