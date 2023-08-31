package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


const (
	host   = "localhost"
	port   = 5432
	user   = "admin"
	dbname = "Avito_segmentations"
)

var db *sql.DB
// Инициализация БД
func Init() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlPassword := os.Getenv("PSQLPass")

	psqlInfo := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable", user, psqlPassword, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
