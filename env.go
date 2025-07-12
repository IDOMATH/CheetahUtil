package env

import (
	"errors"
	"os"
)

func GetEnvValueOrDefault(variable, def string) string {
	str := os.Getenv(variable)
	if str == "" {
		return def
	}
	return str
}

func GetEnvValue(variable string) (string, error) {
	str := os.Getenv(variable)
	if str == "" {
		return "", errors.New("error getting ENV value")
	}
	return str, nil
}
