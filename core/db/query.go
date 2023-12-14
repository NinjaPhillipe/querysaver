package db

import (
	"fmt"
)

// tags

func (sqliteDb *SqliteDb) SelectAllTags() []Tag {
	rows, err := sqliteDb.db.Query("SELECT * FROM tag")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	return readRowTags(rows)
}

func (sqliteDb *SqliteDb) SelectTag(tagId int) (*Tag, error) {
	row, err := sqliteDb.db.Query("SELECT * FROM tag WHERE tag_id = ?", tagId)
	if err != nil {
		// panic(err)
		return nil, err
	}
	if !row.Next() {
		err := fmt.Errorf("No tag with id %d", tagId)
		return nil, err
	}
	return readSingleRowTag(row), nil
}

// todo make color hexadicimal
func (sqliteDb *SqliteDb) InsertTag(name string, color string) {
	fmt.Println("Inserting tag")
	_, err := sqliteDb.db.Exec("INSERT INTO tag (name, color) VALUES (?, ?)", name, color)
	if err != nil {
		panic(err)
	}
}

// folder

func (sqliteDb *SqliteDb) AddFolder(path string) error {
	_, err := sqliteDb.db.Exec("INSERT INTO folder (folder_path) VALUES (?)", path)
	if err != nil {
		return err
	}

	return nil
}

// file

func (sqliteDb *SqliteDb) AddFile(fileName string, folderId int, label string) error {
	_, err := sqliteDb.db.Exec("INSERT INTO file (file_name, fk_folder_id, label) VALUES (?, ?, ?)", fileName, folderId, label)

	if err != nil {
		return err
	}

	return nil
}

// l_file_tag

func (sqliteDb *SqliteDb) AddTagToFile(fileId int, tagId int) error {
	_, err := sqliteDb.db.Exec("INSERT INTO l_file_tag (fk_file_id, fk_tag_id) VALUES (?, ?)", fileId, tagId)

	if err != nil {
		return err
	}

	return nil
}

func (sqliteDb *SqliteDb) RemoveTagFromFile(fileId int, tagId int) error {
	_, err := sqliteDb.db.Exec("DELETE FROM l_file_tag WHERE fk_file_id = ? AND fk_tag_id = ?", fileId, tagId)

	if err != nil {
		return err
	}

	return nil
}
