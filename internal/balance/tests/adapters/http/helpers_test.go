package http_test

import "encoding/json"

func rjsonStr(res string) string {
	return `{"result":` + res + `}`
}

func ejsonStr(res string) string {
	return `{"error":"` + res + `"}`
}

func rjson(res string) string {
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
