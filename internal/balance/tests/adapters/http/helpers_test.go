package http_test

import "encoding/json"

func rjsonStr(res string) string {
	return `{"result":` + res + `}`
}

func ejsonStr(res string) string {
	return `{"error":"` + res + `"}`
}

func rjson(res map[string]any) string {
	json, err := json.Marshal(map[string]any{
		"result": res,
	})
	if err != nil {
		return ""
	}
	return string(json)
}

func ejson(res string) string {
	json, err := json.Marshal(map[string]any{
		"error": res,
	})
	if err != nil {
		return ""
	}
	return string(json)
}

func returnNilError() []any {
	return []any{nil}
}

func returnError(err error) []any {
	return []any{err}
}

func returnSuccess2[T any](out T) []any {
	return []any{out, nil}
}

func returnError2[T any](err error) []any {
	var zero T
	return []any{zero, err}
}
