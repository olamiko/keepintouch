package dbutils

import (
	"database/sql"
	//"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	location string
	name     string
	link     *sql.DB
}

func (db *DB) Init(location string, name string) (*DB, error) {
	var err error
	db.link, err = sql.Open("sqlite3", location)
	if err != nil {
		return nil, err
	}
	db.name = name

	return db, nil
}
