package database

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
