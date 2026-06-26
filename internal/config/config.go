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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		PORT:      os.Getenv("PORT"),
		Dsn:       os.Getenv("DSN"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
