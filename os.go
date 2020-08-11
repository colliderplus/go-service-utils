package service_utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	if value, ok := getenvBool(key); ok {
		return value
	}
	return fallback
}


func getenvBool(key string) (bool, bool) {
	if value, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseBool(value)
		if err != nil {
			return v, true
		}
	}
	return false, false
}