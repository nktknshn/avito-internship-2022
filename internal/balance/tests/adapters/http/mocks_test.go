package http_test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MuxVarsGetterMock struct {
	mock.Mock
}

func (m *MuxVarsGetterMock) GetVars(r *http.Request) map[string]string {
	args := m.Called(r)
	return args.Get(0).(map[string]string)
}

func NewMuxVarsGetterMock(vars map[string]string) *MuxVarsGetterMock {
	m := &MuxVarsGetterMock{}
	m.On("GetVars", mock.Anything).Return(vars)
	return m
}
