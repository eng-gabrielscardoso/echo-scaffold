package config

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func LoadServer() {
	e := echo.New()

	applicationPort := os.Getenv("APP_PORT")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(applicationPort))
}
