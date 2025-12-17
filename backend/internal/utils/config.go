package utils

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func GetEnv(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}

func EnvMust(key string, log *log.Logger) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(MissingRequiredENV.Error(), zap.String("key", key))
	}
	return value
}
