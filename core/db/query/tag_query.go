package query

import (
	. "core/db/data"
	. "core/db/reader"
	"database/sql"
	"fmt"
)

func SelectTag(db *sql.DB, tagId int) (*Tag, error) {
	row, err := db.Query("SELECT * FROM tag WHERE tag_id = ?", tagId)
	if err != nil {
		// panic(err)
		return nil, err
	}
	if !row.Next() {
		err := fmt.Errorf("No tag with id %d", tagId)
		return nil, err
	}
	return ReadSingleRowTag(row), nil
}

// todo make color hexadicimal
func InsertTag(db *sql.DB, name string, color string) error {
	fmt.Println("Inserting tag")
	_, err := db.Exec("INSERT INTO tag (name, color) VALUES (?, ?)", name, color)
	if err != nil { // dumb code
		return err
	}
	return nil
}

func SelectAllTags(db *sql.DB) []Tag {
	rows, err := db.Query("SELECT * FROM tag")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return ReadRowTags(rows)
}
