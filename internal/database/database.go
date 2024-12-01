package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	conn *sql.DB
}

func NewDatabase() (*Database, error) {
	file := "database.db"
	conn, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}
	database := &Database{
		conn: conn,
	}

	return database, nil
}

// read view count (category, slug)
// read reactions (category, slug)

func (db *Database) Buu() {
	type Row struct {
		id   int
		text string
	}

	rows, err := db.conn.Query("select * from greeting")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		r := Row{}
		err := rows.Scan(&r.id, &r.text)
		if err != nil {
			panic(err)
		}
		fmt.Println(r.text)
	}
}
