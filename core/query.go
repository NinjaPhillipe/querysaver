package db

import (
	"fmt"
)

func (sqliteDb *SqliteDb) SelectAllTags() []Tag {
	rows, err := sqliteDb.db.Query("SELECT * FROM tags")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return readRowTag(rows)
}

func (sqliteDb *SqliteDb) SelectTag(tagId int) (Tag, error) {
	row, err := sqliteDb.db.Query("SELECT * FROM tags WHERE tag_id = ?", tagId)
	if err != nil {
		// panic(err)
		return nil, err
	}
	if !row.Next() {
		err := fmt.Errorf("No tag with id %d", tagId)
		return nil, err
	}
	return readRowTag(row), nil
}

// func (sqliteDb *SqliteDb) InsertTag(name string) {
// 	_, err := sqliteDb.db.Exec("INSERT INTO tags  (name) VALUES ('Important')", name)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func (sqliteDb *SqliteDb) InsertTag(name string, color string) {
	fmt.Println("Inserting tag")
	_, err := sqliteDb.db.Exec("INSERT INTO tags  (name, color) VALUES (?, ?)", name, color)
	if err != nil {
		panic(err)
	}
}
