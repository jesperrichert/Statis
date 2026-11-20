package util

import "os"

func Env(key string) (string, bool) {
	env, ok := os.LookupEnv("STATIS_" + key)
	if !ok {
		return "", false
	}
	return env, true
}

func HasEnv(key string) bool {
	env, ok := os.LookupEnv("STATIS_" + key)
	return ok && env != ""
}

func EnvString(key string) string {
	env, ok := os.LookupEnv("STATIS_" + key)
	if !ok {
		return ""
	}
	return env
}
