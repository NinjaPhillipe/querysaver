package query

import "database/sql"

func AddTagToFile(db *sql.DB, fileId int, tagId int) error {
	_, err := db.Exec("INSERT INTO l_file_tag (fk_file_id, fk_tag_id) VALUES (?, ?)", fileId, tagId)

	return err
}

func RemoveTagFromFile(db *sql.DB, fileId int, tagId int) error {
	_, err := db.Exec("DELETE FROM l_file_tag WHERE fk_file_id = ? AND fk_tag_id = ?", fileId, tagId)

	return err
}
