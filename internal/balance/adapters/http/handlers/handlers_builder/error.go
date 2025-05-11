package handlers_builder

type Error struct {
	Error string `json:"error"`
}

func makeErrorBody(err error) any {

	if err == nil {
		return Error{Error: "nil error"}
	}

	return Error{Error: err.Error()}
}
