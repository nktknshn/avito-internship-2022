package server

import (
	"net/http"

	_ "github.com/nktknshn/avito-internship-2022/api/openapi"

	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func swaggerHandler(cfg *config.Config) http.Handler {
	handler := httpSwagger.Handler(
		httpSwagger.PersistAuthorization(true),
		// httpSwagger.URL("swagger.json"),
	)
	return handler
}
