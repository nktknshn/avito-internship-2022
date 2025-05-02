package handlers_builder

type Result[T any] struct {
	Result T `json:"result"`
}

type ResultEmpty struct {
	Result struct{} `json:"result"`
}
