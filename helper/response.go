package helper

func ReponsFormat(code int, msg string, data any) (int, map[string]any) {
	res := map[string]any{}
	res["code"] = code
	res["message"] = msg
	if data != nil {
		res["data"] = data
	}

	return code, res
}
