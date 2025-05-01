package handlers_auth

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/nktknshn/avito-internship-2022/internal/balance/app/use_cases/auth_validate_token"
	ergo "github.com/nktknshn/go-ergo-handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) Handle(ctx context.Context, in auth_validate_token.In) (auth_validate_token.Out, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(auth_validate_token.Out), args.Error(1)
}

func TestUseCaseToValidateToken_EmptyToken(t *testing.T) {
	useCase := &AuthUseCaseMock{}
	validator := &UseCaseToValidateToken{
		useCase: useCase,
	}

	out, ok, err := validator.ValidateToken(context.Background(), "")

	assert.Nil(t, out)
	assert.False(t, ok)
	errWithStatus := err.(ergo.ErrorWithHttpStatus)
	assert.Equal(t, http.StatusBadRequest, errWithStatus.HttpStatusCode)
}

func TestUseCaseToValidateToken_InvalidToken(t *testing.T) {
	useCase := &AuthUseCaseMock{}
	validator := &UseCaseToValidateToken{
		useCase: useCase,
	}

	useCase.On("Handle", mock.Anything, mock.Anything).Return(auth_validate_token.Out{}, auth_validate_token.ErrInvalidToken)

	out, ok, err := validator.ValidateToken(context.Background(), "invalid_token")

	assert.Nil(t, out)
	assert.False(t, ok)
	errWithStatus := err.(ergo.ErrorWithHttpStatus)
	assert.Equal(t, http.StatusBadRequest, errWithStatus.HttpStatusCode)
}

func TestUseCaseToValidateToken_TokenExpired(t *testing.T) {
	useCase := &AuthUseCaseMock{}
	validator := &UseCaseToValidateToken{
		useCase: useCase,
	}

	useCase.On("Handle", mock.Anything, mock.Anything).Return(auth_validate_token.Out{}, auth_validate_token.ErrTokenExpired)

	out, ok, err := validator.ValidateToken(context.Background(), "invalid_token")

	assert.Nil(t, out)
	assert.False(t, ok)
	errWithStatus := err.(ergo.ErrorWithHttpStatus)
	assert.Equal(t, http.StatusUnauthorized, errWithStatus.HttpStatusCode)
}

func TestUseCaseToValidateToken_OtherError(t *testing.T) {
	useCase := &AuthUseCaseMock{}
	validator := &UseCaseToValidateToken{
		useCase: useCase,
	}

	useCase.On("Handle", mock.Anything, mock.Anything).Return(auth_validate_token.Out{}, errors.New("some internal error"))

	out, ok, err := validator.ValidateToken(context.Background(), "invalid_token")

	assert.Nil(t, out)
	assert.False(t, ok)
	assert.Error(t, err)
	assert.Equal(t, "some internal error", err.Error())
}

func TestUseCaseToValidateToken_Success(t *testing.T) {
	useCase := &AuthUseCaseMock{}
	validator := &UseCaseToValidateToken{
		useCase: useCase,
	}

	useCase.On("Handle", mock.Anything, mock.Anything).Return(auth_validate_token.Out{}, nil)

	out, ok, err := validator.ValidateToken(context.Background(), "valid_token")

	assert.NotNil(t, out)
	assert.True(t, ok)
	assert.NoError(t, err)
}
