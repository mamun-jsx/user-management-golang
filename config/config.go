package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Declare a struct of config
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

// make function to call into anywhere
func LoadEnv() *Config {
	//* call dodotenv.Load() function from package
	if err := godotenv.Load(); err != nil {
		log.Fatal("env variable required")
	}
	// return and store all env values into config
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DbHost:     getEnv("DB_HOST", "localhost"),
		DbPort:     getEnv("DB_PORT", "5432"),
		DbUser:     getEnv("DB_USER", "postgres"),
		DbPassword: getEnv("DB_PASSWORD", "mysecretpassword"),
		DbName:     getEnv("DB_NAME", "user_management_app"),
	}
}
