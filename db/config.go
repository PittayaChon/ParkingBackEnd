package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	POSTGRES_HOST     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	POSTGRES_PORT     string
}

func LoadENV() env {
	if os.Getenv("GO_ENV") != "production" {

		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	loadedEnv := env{
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DB:       os.Getenv("POSTGRES_DB"),
		POSTGRES_PORT:     os.Getenv("POSTGRES_PORT"),
	}

	return loadedEnv
}
