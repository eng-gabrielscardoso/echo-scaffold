package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDatabase() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dns := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=%s TimeZone=%s",
		username, password, database, port, host, "disable", "Asia/Kolkata")

	db, error := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	}

	connection, error := db.DB()

	if error != nil {
		log.Fatal(error)
	}

	return connection
}
