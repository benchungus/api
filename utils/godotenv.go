package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnvVar() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("error loading .env file")
	}
}
