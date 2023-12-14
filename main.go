package main

import (
	"core/db"
	"core/server"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	db.Add(1, 2)

	// sqliteDb := &db.SqliteDb{}
	// sqliteDb.Connect()

	// tags := sqliteDb.SelectAllTags()
	// fmt.Println(tags)
	// tagOne, _ := sqliteDb.SelectTag(2)
	// fmt.Println(tagOne)
	// // sqliteDb.InsertTag("Yellow", "#FF0000")

	serv := server.NewServer()
	serv.Init()
	serv.Run()

	// serv.Close()

	// sqliteDb.Close()
}
