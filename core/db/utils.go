package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const DEFAULT_ARRAY_SIZE = 1000

type Tag struct {
	Id    int
	Name  string
	Color string
}

type File struct {
	Id       int
	Name     string
	FolderId int
	Label    string
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

func readSingleRowFile(rows *sql.Rows) *File {
	file := File{}
	if err := rows.Scan(&file.Id, &file.Name, &file.FolderId, &file.Label); err != nil {
		log.Fatal(err)
	}
	return &file
}

func readRowFiles(rows *sql.Rows) []File {
	res := make([]File, 0, DEFAULT_ARRAY_SIZE)
	for rows.Next() {
		file := readSingleRowFile(rows)
		res = append(res, *file)
	}
	fmt.Println(res)
	return res
}

func concatArrayInt(nbs []int, sep string) string {

	strNumbers := make([]string, len(nbs))

	for i, num := range nbs {
		strNumbers[i] = strconv.Itoa(num)
	}
	return strings.Join(strNumbers, sep)
}
