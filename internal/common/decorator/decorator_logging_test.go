package decorator_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/nktknshn/avito-internship-2022/internal/common/decorator"
)

type useCase0Mock struct {
	mock.Mock
}

func (u *useCase0Mock) Handle(ctx context.Context, in int) error {
	args := u.Called(ctx, in)
	return args.Error(0)
}

func (u *useCase0Mock) GetName() string {
	return "useCase0"
}

type useCase1Mock struct {
	mock.Mock
}

func (u *useCase1Mock) Handle(ctx context.Context, in int) (int, error) {
	args := u.Called(ctx, in)
	return args.Get(0).(int), args.Error(1)
}

func (u *useCase1Mock) GetName() string {
	return "useCase1"
}

type loggerMock struct {
	mock.Mock
}

func (l *loggerMock) Info(message string, keyvals ...interface{}) {
	l.Called(message, keyvals)
}

func (l *loggerMock) Error(message string, keyvals ...interface{}) {
	l.Called(message, keyvals)
}

func (l *loggerMock) Debug(message string, keyvals ...interface{}) {
	l.Called(message, keyvals)
}

func (l *loggerMock) Fatal(message string, keyvals ...interface{}) {
	l.Called(message, keyvals)
}

func (l *loggerMock) InitLogger(args ...interface{}) {
	l.Called(args...)
}

func (l *loggerMock) Warn(message string, keyvals ...interface{}) {
	l.Called(message, keyvals)
}

func (l *loggerMock) GetLogger() any {
	return nil
}

func TestDecorator0Logging_Success(t *testing.T) {
	logger := &loggerMock{}
	ucmock := &useCase0Mock{}

	ucmock.On("Handle", mock.Anything, mock.Anything).Return(nil)

	logger.On("Info", mock.Anything, mock.Anything).Return(nil)

	decorator := decorator.Decorator0Logging[int]{
		Base:   ucmock,
		Logger: logger,
	}
	decorator.Handle(t.Context(), 1)

	ucmock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestDecorator0Logging_Error(t *testing.T) {
	logger := &loggerMock{}
	ucmock := &useCase0Mock{}

	ucmock.On("Handle", mock.Anything, mock.Anything).Return(errors.New("error"))
	logger.On("Error", mock.Anything, mock.Anything).Return(nil)
	logger.On("Info", mock.Anything, mock.Anything).Return(nil)

	decorator := decorator.Decorator0Logging[int]{
		Base:   ucmock,
		Logger: logger,
	}
	decorator.Handle(t.Context(), 1)

	ucmock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestDecorator1Logging_Success(t *testing.T) {
	logger := &loggerMock{}
	ucmock := &useCase1Mock{}

	ucmock.On("Handle", mock.Anything, mock.Anything).Return(1, nil)
	logger.On("Info", mock.Anything, mock.Anything).Return(nil)

	decorator := decorator.Decorator1Logging[int, int]{
		Base:   ucmock,
		Logger: logger,
	}
	decorator.Handle(t.Context(), 1)

	ucmock.AssertExpectations(t)
	logger.AssertExpectations(t)
}

func TestDecorator1Logging_Error(t *testing.T) {
	logger := &loggerMock{}
	ucmock := &useCase1Mock{}

	ucmock.On("Handle", mock.Anything, mock.Anything).Return(0, errors.New("error"))
	logger.On("Error", mock.Anything, mock.Anything).Return(nil)
	logger.On("Info", mock.Anything, mock.Anything).Return(nil)

	decorator := decorator.Decorator1Logging[int, int]{
		Base:   ucmock,
		Logger: logger,
	}
	decorator.Handle(t.Context(), 1)

	ucmock.AssertExpectations(t)
	logger.AssertExpectations(t)
}
