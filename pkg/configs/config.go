package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr                   string `env:"ADDR"`
	BasePrefix             string `env:"BASE_PREFIx"`
	HttpPort               string `env:"HTTP_PORT"`
	DBConnectionString     string `env:"DB_CONNECTION_STRINg"`
	JwtSecret              string `env:"JWT_SECRET"`
	JwtExpirationInSeconds int64  `env:"JWT_EXPIRATION_IN_SECONDS"`
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Addr:                   getEnv("ADDR", "localhost"),
		BasePrefix:             getEnv("BASE_PREFIX", "/api/v1"),
		HttpPort:               getEnv("PORT", ":8080"),
		DBConnectionString:     getEnv("DB_CONNECTION_STRING", "https://localhost:5173"),
		JwtSecret:              getEnv("JWT_SCRET", "THIS_A_SUPER_SCRET"),
		JwtExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 600000),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
