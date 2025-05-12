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

func (m *Noop) IncHits(_ Status, _ string) {}

func (m *Noop) ObserveResponseTime(_ Status, _ string, _ float64) {}
