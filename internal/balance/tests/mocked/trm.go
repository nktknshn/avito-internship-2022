package mocked

import (
	"context"

	"github.com/avito-tech/go-transaction-manager/trm"
)

type TrmManagerMock struct {
}

func (m *TrmManagerMock) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

func (m *TrmManagerMock) DoWithSettings(ctx context.Context, settings trm.Settings, fn func(ctx context.Context) error) error {
	return fn(ctx)
}

var _ trm.Manager = &TrmManagerMock{}
