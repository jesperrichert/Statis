package util

func CheckNull(data any) bool {
	if data == nil {
		return true
	}
	str, ok := data.(string)
	if ok {
		if len(str) <= 0 {
			return true
		}
		return false
	}
	return false
}

func CheckNullLen[T any](data []T) bool {
	if data == nil {
		return true
	}
	if len(data) <= 0 {
		return true
	}
	return false
}
