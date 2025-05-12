package server

import (
	"net/http"

	// подключаем OpenAPI спецификацию
	_ "github.com/nktknshn/avito-internship-2022/api/openapi"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

func swaggerHandler(_ *config.Config) http.Handler {
	handler := httpSwagger.Handler(
		httpSwagger.PersistAuthorization(true),
		// httpSwagger.URL("swagger.json"),
	)
	return handler
}
