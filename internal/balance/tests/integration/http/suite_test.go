package http_test

import (
	"testing"

	adaptersHttp "github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http"
	"github.com/nktknshn/avito-internship-2022/internal/balance/tests/mocked"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
	"github.com/stretchr/testify/suite"
)

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}

type HttpTestSuite struct {
	testing_pg.TestSuitePg
	app         *mocked.AppMocked
	httpAdapter *adaptersHttp.HttpAdapter
}

func (s *HttpTestSuite) SetupTest() {
	s.app = mocked.NewMockedApp()

}
