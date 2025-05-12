//nolint:nonamedreturns // используем в defer
package decorator

import (
	"context"
	"time"

	commonErrors "github.com/nktknshn/avito-internship-2022/internal/common/errors"
	"github.com/nktknshn/avito-internship-2022/internal/common/metrics"
)

type Decorator0Metrics[T any] struct {
	Base    UseCase0Handler[T]
	Metrics metrics.Metrics
}

func (d *Decorator0Metrics[T]) Handle(ctx context.Context, in T) (err error) {
	started := time.Now()
	defer func() {
		status := metrics.StatusSuccess
		if err != nil {
			status = metrics.StatusError
		}
		if commonErrors.IsPanicError(err) {
			status = metrics.StatusPanic
		}
		d.Metrics.IncHits(status, d.Base.GetName())
		d.Metrics.ObserveResponseTime(status, d.Base.GetName(), time.Since(started).Seconds())
	}()

	return d.Base.Handle(ctx, in)
}

func (d *Decorator0Metrics[T]) GetName() string {
	return d.Base.GetName()
}

type Decorator1Metrics[T any, R any] struct {
	Base    UseCase1Handler[T, R]
	Metrics metrics.Metrics
}

func (d *Decorator1Metrics[T, R]) Handle(ctx context.Context, in T) (result R, err error) {
	started := time.Now()
	defer func() {
		status := metrics.StatusSuccess
		if err != nil {
			status = metrics.StatusError
		}
		d.Metrics.IncHits(status, d.Base.GetName())
		d.Metrics.ObserveResponseTime(status, d.Base.GetName(), time.Since(started).Seconds())
	}()

	return d.Base.Handle(ctx, in)
}

func (d *Decorator1Metrics[T, R]) GetName() string {
	return d.Base.GetName()
}
