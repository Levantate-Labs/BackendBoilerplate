package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
	// Add config variables here
}

var config *Config

func InitConfig() {
	env := os.Getenv("APP_ENV")
	if env == "DEV" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}

	config = &Config{
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		// Add config variables here
	}
}

func LoadConfig() *Config {
	return config
}
