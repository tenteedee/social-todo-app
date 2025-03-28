package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	DbPort     string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbName = os.Getenv("DB_NAME")
	DbPort = os.Getenv("DB_PORT")
}
