package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const dbsource = "postgresql://<user>:<password>@localhost:<port>/<dbname>?sslmode=disable"

var db *sql.DB

func GetSession() (*sql.DB, error) {
	var err error

	db, err = sql.Open("postgres", dbsource)
	if err != nil {
		return nil, err
	}

	return db, nil
}
