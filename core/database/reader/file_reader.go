package reader

import (
	"core/database/data"
	"database/sql"
	"fmt"
	"log"
)

func ReadSingleRowFile(rows *sql.Rows) *data.File {
	file := data.File{}
	if err := rows.Scan(&file.Id, &file.Name, &file.FolderId, &file.Label); err != nil {
		log.Fatal(err)
	}
	return &file
}

func ReadRowFiles(rows *sql.Rows) []data.File {
	res := make([]data.File, 0, DEFAULT_ARRAY_SIZE)
	for rows.Next() {
		file := ReadSingleRowFile(rows)
		res = append(res, *file)
	}
	fmt.Println(res)
	return res
}
