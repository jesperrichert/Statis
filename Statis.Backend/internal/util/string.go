package util

import "strings"

func ReplaceStrings(str string, strMap map[string]string) string {
	s := str
	for key, value := range strMap {
		s = strings.ReplaceAll(s, key, value)
	}
	return s
}

func ReplacePlaceholders(str string, placeholder map[string]string) string {
	return ReplaceStrings(str, placeholder)
}
