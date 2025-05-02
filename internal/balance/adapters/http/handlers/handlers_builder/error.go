package handlers_builder

type Error struct {
	Error string `json:"error"`
}

func makeErrorBody(err error) any {
	return Error{Error: err.Error()}
}
