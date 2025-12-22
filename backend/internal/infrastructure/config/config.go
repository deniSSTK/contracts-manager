package config

import (
	"contracts-manager/internal/domain/auth"
	"contracts-manager/internal/infrastructure/db/models"
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

	Admin auth.SignupDTO
}

func NewConfig() *Config {
	_ = godotenv.Load()

	log := logger.New(os.Stderr, "[CONFIG] ", logger.LstdFlags)
	log.Print("Loading configuration...")

	admintype := models.UserTypeAdmin

	return &Config{
		Port:    utils.EnvMust("PORT", log),
		RunMode: utils.GetEnv("RUN_MODE", "dev"),

		FrontendUrl: utils.EnvMust("FRONTEND_URL", log),

		JWTSecret: utils.EnvMust("JWT_SECRET", log),

		Admin: auth.SignupDTO{
			Username: utils.GetEnv("ADMIN_USERNAME", "admin_user"),
			Email:    utils.GetEnv("ADMIN_EMAIL", "admin@email.com"),
			Password: utils.GetEnv("ADMIN_PASSWORD", "admin_pass"),
			Type:     &admintype,
		},
	}
}
