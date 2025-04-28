package http_test

import (
	"github.com/nktknshn/avito-internship-2022/internal/balance/app"
	"github.com/nktknshn/avito-internship-2022/pkg/testing_pg"
)

type HttpTestSuite struct {
	testing_pg.TestSuitePg
	app *app.Application
}
