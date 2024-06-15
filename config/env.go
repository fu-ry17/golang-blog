package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PORT        string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_ADDRESS  string
	DB_HOST     string
	DB_NAME     string
}

var Env = initStorage()

func initStorage() Config {
	godotenv.Load(".env")
	return Config{
		DB_NAME:     getEnv("DB_NAME", "blog"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "pass123"),
		DB_USER:     getEnv("DB_USER", "test"),
		DB_ADDRESS: fmt.Sprintf("%s%s",
			getEnv("DB_HOST", "db"),
			getEnv("DB_PORT", ":3306"),
		),
		PORT: getEnv("PORT", ":3001"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
