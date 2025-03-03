package config

import (
	"fmt"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConf struct {
	GoEnv         string `env:"GO_ENV" envDefault:"development"`
	AppHost       string `env:"APP_HOST" envDefault:"localhost"`
	AppPort       int    `env:"APP_PORT" envDefault:"3300"`
	DBHost        string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort        int    `env:"DB_PORT" envDefault:"3306"`
	DBUsername    string `env:"DB_USERNAME"`
	DBPassword    string `env:"DB_PASSWORD"`
	DBName        string `env:"DB_NAME"`
	RedisHost     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort     int    `env:"REDIS_PORT" envDefault:"6379"`
	RedisUsername string `env:"REDIS_USERNAME"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

var Env EnvConf

func LoadEnv() error {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "" {
		goEnv = "development"
	}

	envFile := ".env." + goEnv

	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("failed to load env file: %w", err)
	}

	if err := env.Parse(&Env); err != nil {
		return fmt.Errorf("error loading environment variables: %w", err)
	}

	return nil

}
