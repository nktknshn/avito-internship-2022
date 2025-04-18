package decorator

import (
	"context"
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
)

type DecoratorCommandMetrics[T any] struct {
	base       UseCaseCommandHandler[T]
	metrics    metrics.Metrics
	methodName string
}

func (d *DecoratorCommandMetrics[T]) Handle(ctx context.Context, in T) error {
	return d.base.Handle(ctx, in)
}

type DecoratorQueryMetrics[T any, R any] struct {
	base       UseCaseQueryHandler[T, R]
	m          metrics.Metrics
	methodName string
}

func (d *DecoratorQueryMetrics[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	started := time.Now()
	defer func() {
		status := metrics.StatusSuccess
		if err != nil {
			status = metrics.StatusError
		}
		d.m.IncHits(status, d.methodName)
		d.m.ObserveResponseTime(status, d.methodName, time.Since(started).Seconds())
	}()

	return d.base.Handle(ctx, in)
}
