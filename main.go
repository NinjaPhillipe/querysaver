package main

import (
	db "core"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	db.Add(1, 2)

	sqliteDb := &db.SqliteDb{}
	sqliteDb.Connect()

	tags := sqliteDb.SelectAllTags()
	fmt.Println(tags)
	// sqliteDb.SelectTag(1)
	// sqliteDb.InsertTag("Yellow", "#FF0000")

	sqliteDb.Close()
}
