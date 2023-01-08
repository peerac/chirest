package config

import (
	"log"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port     int
	Env      string
	Version  string
	Database struct {
		DSN string
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
