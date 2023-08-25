package db

import (
	"database/sql"
	"fmt"
)

type Segment struct {
	userID  int      `json:"user_id"`
	Items []string `json:"Items"`
}

type Segmentations struct {
	Segments	[]Segment
}
const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "76705315Rr"
	dbname   = "Avito_segmentations"
)

var db *sql.DB

func Init() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s?sslmode=disable", user, password, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
