package databasecontrol

import (
	"database/sql"
)

func SelectFromDB(db *sql.DB) {
	query := "SELECT id, name, note FROM notes"
	rows, _ := db.Query(query)
	var id int
	var name string
	var note string

	for rows.Next() {
		rows.Scan(&id, &name, &note)
	}
}
