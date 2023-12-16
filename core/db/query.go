package db

// folder

func (sqliteDb *SqliteDb) AddFolder(path string) error {
	_, err := sqliteDb.Db.Exec("INSERT INTO folder (folder_path) VALUES (?)", path)
	if err != nil {
		return err
	}

	return nil
}

// file

func (sqliteDb *SqliteDb) AddFile(fileName string, folderId int, label string) error {
	_, err := sqliteDb.Db.Exec("INSERT INTO file (file_name, fk_folder_id, label) VALUES (?, ?, ?)", fileName, folderId, label)

	if err != nil {
		return err
	}

	return nil
}

func (SqliteDb *SqliteDb) QueryFileWithTagsIdOr(tagsId []int) (*[]File, error) {
	tags := concatArrayInt(tagsId, ",")

	row, err := SqliteDb.Db.Query(`select f.*
	FROM file f
	join l_file_tag lft on lft.fk_file_id  = f.id_file
	where lft.fk_tag_id  = 1`, tags)

	if err != nil {
		return nil, err
	}

	files := readRowFiles(row)

	return &files, nil
}
