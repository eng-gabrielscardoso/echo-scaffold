package internal

import (
	config "github.com/eng-gabrielscardoso/echo-scaffold/internal/config"
)

// This function bootstraps the application
func Bootstrap() {
	config.LoadDotenv()
	connection := config.LoadDatabase()
	config.LoadServer()

	defer connection.Close()
}
