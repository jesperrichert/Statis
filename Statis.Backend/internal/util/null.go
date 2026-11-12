package util

func CheckNull(data any) bool {
	if data == nil {
		return true
	}
	return false
}
