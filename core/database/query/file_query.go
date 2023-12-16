package query

import (
	"core/database/data"
	"core/database/reader"
	"database/sql"
	"fmt"
	"strings"
)

func QueryFileWithTagsIdOr(db *sql.DB, tagsId []int) (*[]data.File, error) {
	// Generate placeholders for the values
	placeholders := make([]string, len(tagsId))
	for i := range placeholders {
		placeholders[i] = "?"
	}

	// Generate the query
	query := fmt.Sprintf(`SELECT f.*
	FROM file f
	JOIN l_file_tag lft ON lft.fk_file_id = f.id_file
	WHERE lft.fk_tag_id IN (%s)`, strings.Join(placeholders, ","))

	// Convert tagsId to []interface{}
	args := make([]interface{}, len(tagsId))
	for i, v := range tagsId {
		args[i] = v
	}

	// Execute the query
	row, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	files := reader.ReadRowFiles(row)

	return &files, nil
}
