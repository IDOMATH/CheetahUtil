package env

import (
	"os"
)

func GetEnvValueOrDefault(variable, def string) string {
	str := os.Getenv(variable)
	if str == "" {
		return def
	}
	return str
}

func GetEnvValue(variable string) string {
	return os.Getenv(variable)
}
