package metrics

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type Metrics interface {
	IncHits(status Status, method string)
	ObserveResponseTime(status Status, method string, observeTime float64)
}
