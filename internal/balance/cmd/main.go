package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/handlers"
	"github.com/nktknshn/avito-internship-2022/internal/balance/adapters/http/router"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/nktknshn/avito-internship-2022/internal/balance/service"
)

var (
	flagConfigPath = "config.yaml"
)

func main() {

	flag.StringVar(&flagConfigPath, "cfg_path", flagConfigPath, "config path")
	flag.Parse()

	ctx := context.Background()
	cfg := config.Config{}
	err := cleanenv.ReadConfig(flagConfigPath, &cfg)

	if err != nil {
		panic(err)
	}

	app, err := service.NewApplication(ctx, &cfg)

	if err != nil {
		panic(err)
	}

	handlers := handlers.CreateHandlers(app)
	router := router.CreateMuxRouter(handlers)

	mux := http.NewServeMux()
	mux.Handle("/", router)

	http.ListenAndServe(":8080", mux)

	_ = app
}
