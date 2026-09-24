package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

/**
# Server Configuration
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=user_management_app
**/
//Declare a struct of config
type Config struct {
	AppPort    string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

// create func to take env value from .env file
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env variable required")
	}

	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbPassword: getEnv("DB_PASSWORD", "mysecretpassword"),
		DbName: getEnv("DB_NAME","user_management_app"),
	}
}
