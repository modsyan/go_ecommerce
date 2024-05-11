package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort               string `env:"http_port"`
	DBConnectionString     string `env:"db_connection_string"`
	JwtSecret              string `env:"jwt_secret"`
	JwtExpirationInSeconds int64  `env:"jwt_expiration_in_seconds"`
}

var config = initConfig()

func initConfig() Config {

	return Config {
		HttpPort: getEnv()
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getConfig() Config {

	return config
}
