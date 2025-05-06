package decorator

import (
	"context"
	"time"

	commonErrors "github.com/nktknshn/avito-internship-2022/internal/common/errors"
	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
)

type Decorator0Metrics[T any] struct {
	base    UseCase0Handler[T]
	metrics metrics.Metrics
}

func (d *Decorator0Metrics[T]) Handle(ctx context.Context, in T) (err error) {
	started := time.Now()
	defer func() {
		status := metrics.StatusSuccess
		if err != nil {
			status = metrics.StatusError
		}
		if commonErrors.IsErrPanic(err) {
			status = metrics.StatusPanic
		}
		d.metrics.IncHits(status, d.base.GetName())
		d.metrics.ObserveResponseTime(status, d.base.GetName(), time.Since(started).Seconds())
	}()

	return d.base.Handle(ctx, in)
}

func (d *Decorator0Metrics[T]) GetName() string {
	return d.base.GetName()
}

type Decorator1Metrics[T any, R any] struct {
	base    UseCase1Handler[T, R]
	metrics metrics.Metrics
}

func (d *Decorator1Metrics[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
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

func (d *Decorator1Metrics[T, R]) GetName() string {
	return d.base.GetName()
}
