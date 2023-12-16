package reader

import (
	. "core/database/data"
	"database/sql"
	"fmt"
	"log"
)

func ReadSingleRowTag(rows *sql.Rows) *Tag {
	tag := Tag{}
	if err := rows.Scan(&tag.Id, &tag.Name, &tag.Color); err != nil {
		log.Fatal(err)
	}
	return &tag
}

func ReadRowTags(rows *sql.Rows) []Tag {
	res := make([]Tag, 0, DEFAULT_ARRAY_SIZE)
	for rows.Next() {
		tag := ReadSingleRowTag(rows)
		res = append(res, *tag)
	}
	fmt.Println(res)
	return res
}
