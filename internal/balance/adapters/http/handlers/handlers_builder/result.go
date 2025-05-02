package handlers_builder

type Result[T any] struct {
	Result T `json:"result"`
}

type ResultEmpty struct {
	Result struct{} `json:"result"`
}

func makeResultBody(result any) any {

	if result == nil {
		return ResultEmpty{}
	}

	return Result[any]{Result: result}
}
