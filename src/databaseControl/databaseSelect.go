package databasecontrol

import (
	"database/sql"
	"log"
)

type Test struct {
	ID   int
	Name string
	Note string
}

// selecting from db all data.
func SelectFromDBrow(db *sql.DB) {
	var test Test
	query := "SELECT id, name, note FROM notes"
	rows, _ := db.Query(query)
	var id int
	var name string
	var note string

	for rows.Next() {
		rows.Scan(&id, &name, &note)
		test.ID = id
		test.Name = name
		test.Note = note
		log.Println("ID: ", test.ID, "Name: ", test.Name, "Note:", test.Note)
	}
}

// selecting from db by name.
func SelectFromDBbyName() {
	return
}
