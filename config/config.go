package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APPport    string
	DBhost     string
	DBport     string
	DBuser     string
	DBpassword string
	DBname     string
}

// get env key from .env file
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// load env from config

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can not find env variables")
	}
	return &Config{
		APPport:    getEnv("APP_PORT", "8080"),
		DBhost:     getEnv("DB_HOST", "localhost"),
		DBport:     getEnv("DB_PORT", "5432"),
		DBuser:     getEnv("DB_USER", "postgres"),
		DBpassword: getEnv("DB_PASSWORD", "mysecretpassword"),
		DBname:     getEnv("DB_NAME", "user_management_app"),
	}
}
