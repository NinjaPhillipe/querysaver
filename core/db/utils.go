package db

import (
	"database/sql"
	"fmt"
	"log"
)

const DEFAULT_ARRAY_SIZE = 1000

type Tag struct {
	Id    int
	Name  string
	Color string
}

func readSingleRowTag(rows *sql.Rows) *Tag {
	tag := Tag{}
	if err := rows.Scan(&tag.Id, &tag.Name, &tag.Color); err != nil {
		log.Fatal(err)
	}
	return &tag
}

func readRowTags(rows *sql.Rows) []Tag {
	res := make([]Tag, 0, DEFAULT_ARRAY_SIZE)
	for rows.Next() {
		tag := readSingleRowTag(rows)
		res = append(res, *tag)
	}
	fmt.Println(res)
	return res
}