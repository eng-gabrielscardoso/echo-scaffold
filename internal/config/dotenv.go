package config

import (
	"log"

	"github.com/joho/godotenv"
)

// This function loads the `.env` file using `joho/godotenv`
func LoadDotenv() {
	if error := godotenv.Load(); error != nil {
		log.Fatal(error)
	}
}
