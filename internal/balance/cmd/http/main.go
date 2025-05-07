package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nktknshn/avito-internship-2022/internal/balance/cmd/http/server"
	"github.com/nktknshn/avito-internship-2022/internal/balance/config"
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jaegermetrics "github.com/uber/jaeger-lib/metrics"
)

var (
	flagConfigPath = "config.yaml"
)

func main() {

	flag.StringVar(&flagConfigPath, "cfg_path", flagConfigPath, "config path")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.LoadConfigFromFile(flagConfigPath)

	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           cfg.Jaeger.LogSpans,
			LocalAgentHostPort: cfg.Jaeger.Host,
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(jaegermetrics.NullFactory),
	)

	if err != nil {
		log.Fatalf("NewTracer: %v", err)
	}

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	server := server.NewHttpServer(cfg)

	if err := server.Init(ctx); err != nil {
		panic(err)
	}

	err = server.Run(ctx)
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	err = server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}
