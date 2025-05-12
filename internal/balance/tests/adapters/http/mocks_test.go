package http_test

import (
	"net/http"
)

type MuxVarsGetterMock struct {
	vars map[string]string
}

func NewMuxVarsGetterMock(vars map[string]string) *MuxVarsGetterMock {
	return &MuxVarsGetterMock{vars: vars}
}

func (m *MuxVarsGetterMock) GetVar(_ *http.Request, key string) (string, bool) {
	return m.vars[key], m.vars[key] != ""
}
