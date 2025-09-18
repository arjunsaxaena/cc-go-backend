package config

import (
	"microservice/common/environment"
	"microservice/common/logger"
)

type Config struct {
	DATABASE_URL string `env:"DATABASE_URL"`
	DATABASE_URL_NEON string `env:"DATABASE_URL_NEON"`
	SECRET_KEY string `env:"SECRET_KEY"`
	UPLOAD_DIR string `env:"UPLOAD_DIR"`
	GEMINI_API_KEY string `env:"GEMINI_API_KEY"`
	HUGGINGFACE_TOKEN string `env:"HUGGING_FACE_TOKEN"`
}

func LoadConfig() *Config {
	config := Config{}
	if err := environment.LoadEnv(&config); err != nil {
		logger.Panic("unable to load env: %s", err)
		return nil
	}
	return &config
}
