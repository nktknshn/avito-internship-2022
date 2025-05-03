package decorator

import (
	"context"
	"time"

	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
)

type DecoratorCommandMetrics[T any] struct {
	base    UseCaseCommandHandler[T]
	metrics metrics.Metrics
}

func (d *DecoratorCommandMetrics[T]) Handle(ctx context.Context, in T) (err error) {
	started := time.Now()
	defer func() {
		status := metrics.StatusSuccess
		if err != nil {
			status = metrics.StatusError
		}
		d.metrics.IncHits(status, d.base.GetName())
		d.metrics.ObserveResponseTime(status, d.base.GetName(), time.Since(started).Seconds())
	}()

	return d.base.Handle(ctx, in)
}

func (d *DecoratorCommandMetrics[T]) GetName() string {
	return d.base.GetName()
}

type DecoratorQueryMetrics[T any, R any] struct {
	base    UseCaseQueryHandler[T, R]
	metrics metrics.Metrics
}

func (d *DecoratorQueryMetrics[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	started := time.Now()
	defer func() {
		status := metrics.StatusSuccess
		if err != nil {
			status = metrics.StatusError
		}
		d.metrics.IncHits(status, d.base.GetName())
		d.metrics.ObserveResponseTime(status, d.base.GetName(), time.Since(started).Seconds())
	}()

	return d.base.Handle(ctx, in)
}

func (d *DecoratorQueryMetrics[T, R]) GetName() string {
	return d.base.GetName()
}
