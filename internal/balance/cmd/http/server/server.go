package server

import "github.com/nktknshn/avito-internship-2022/internal/balance/config"

type HttpServer struct {
}

func NewHttpServer(cfg *config.ConfigHTTP) *HttpServer {
	return &HttpServer{}
}
