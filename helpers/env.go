package helpers

import (
	"os"
	"strconv"
)

func GetEnvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	var valueInt int
	var err error

	if valueInt, err = strconv.Atoi(value); err != nil {
		return defaultValue
	}

	return valueInt
}