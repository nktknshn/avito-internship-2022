package main

import (
	"net/http"

	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
)

func setTLS(server *http.Server, cfg *config.ConfigTLS) {
	// server.TLSConfig = &tls.Config{
	// 	Certificates: []tls.Certificate{cfg.GetCertificates()},
	// }
}
