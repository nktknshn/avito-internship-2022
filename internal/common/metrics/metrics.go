package metrics

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
	StatusPanic   Status = "panic"
)

func (s Status) String() string {
	return string(s)
}

type Metrics interface {
	IncHits(status Status, method string)
	ObserveResponseTime(status Status, method string, observeTime float64)
}

type Noop struct{}

func (m *Noop) IncHits(status Status, method string) {}

func (m *Noop) ObserveResponseTime(status Status, method string, observeTime float64) {}
