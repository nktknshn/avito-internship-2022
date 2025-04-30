package http_test

func rjson(res string) string {
	return `{"result":` + res + `}`
}

func ejson(res string) string {
	return `{"error":"` + res + `"}`
}
