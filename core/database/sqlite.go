package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO REFACTO REMOVE  SqliteDb struct

type SqliteDb struct {
	Db *sql.DB
}

func (sqliteDb *SqliteDb) Connect() {
	fmt.Println("Sqlite start connecting")

	var err error

	sqliteDb.Db, err = sql.Open("sqlite3", "./build/querysaverdb.db")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sqlite connected")
}

func (sqliteDb *SqliteDb) Close() {
	fmt.Println("Sqlite closed")
	defer sqliteDb.Db.Close()
}
