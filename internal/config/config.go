package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	Dsn       string
	JwtSecret string
}

func LoadEnv() *Config {

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using environment variables")
	}

	return &Config{
		PORT:      os.Getenv("PORT"),
		Dsn:       os.Getenv("DSN"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
