package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	db *sql.DB
}

func (sqliteDb *SqliteDb) Connect() {
	fmt.Println("Sqlite start connecting")

	var err error

	sqliteDb.db, err = sql.Open("sqlite3", "./build/querysaverdb.db")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sqlite connected")
}

func (sqliteDb *SqliteDb) Close() {
	fmt.Println("Sqlite closed")
	defer sqliteDb.db.Close()
}
