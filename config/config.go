package config

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppServiceName string `env:"APP_SERVICE_NAME"`
	AppEnv         string `env:"APP_ENV"`
	AppVersion     string `env:"APP_VERSION"`

	LogLevel string `env:"LOG_LEVEL"`

	HttpPort         string `env:"HTTP_PORT"`
	HttpReadTimeout  int    `env:"HTTP_READ_TIMEOUT"`
	HttpWriteTimeout int    `env:"HTTP_WRITE_TIMEOUT"`
}

func LoadEnvConfig() Config {
	var config Config

	godotenv.Load(".env")
	ctx := context.Background()
	if err := envconfig.Process(ctx, &config); err != nil {
		panic(err)
	}

	return config
}
