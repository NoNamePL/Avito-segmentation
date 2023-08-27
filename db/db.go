package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Segment struct {
	userID int      `json:"user_id"`
	Items  []string `json:"Items"`
}

type Segmentations struct {
	Segments []Segment
}

const (
	host   = "localhost"
	port   = 5432
	user   = "admin"
	dbname = "Avito_segmentations"
)

var db *sql.DB

func Init() (*sql.DB, error) {

	psqlPassword, ok := os.LookupEnv("PSQLPass")
	if !ok {
		log.Fatal("Can't connect to .env")
	}
	psqlInfo := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable", user, psqlPassword, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
