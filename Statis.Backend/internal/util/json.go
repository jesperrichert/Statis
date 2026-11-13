package util

import (
	"bytes"
	"encoding/json"
	"reflect"
)

func CanUnmarshalAs[T any](data []byte) bool {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	var d T
	return dec.Decode(&d) == nil
}

func IsEmpty(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map, reflect.Slice:
		return rv.Len() == 0
	default:
		return rv.IsZero()
	}
}
