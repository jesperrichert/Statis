package util

func JsonResult(message string, data map[string]any) map[string]any {
	var dto map[string]any = map[string]any{}
	dto["message"] = message

	for k, v := range data {
		dto[k] = v
	}

	return dto
}
