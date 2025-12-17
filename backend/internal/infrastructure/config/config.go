package config

import (
	"contracts-manager/internal/utils"
	logger "log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	RunMode string

	FrontendUrl string

	JWTSecret string
}

func NewConfig() *Config {
	_ = godotenv.Load()

	log := logger.New(os.Stderr, "[CONFIG] ", logger.LstdFlags)
	log.Print("Loading configuration...")

	return &Config{
		Port:    utils.EnvMust("PORT", log),
		RunMode: utils.GetEnv("RUN_MODE", "dev"),

		FrontendUrl: utils.EnvMust("FRONTEND_URL", log),

		JWTSecret: utils.EnvMust("JWT_SECRET", log),
	}
}
