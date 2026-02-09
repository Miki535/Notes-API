package databasecontrol

import (
	"database/sql"
	"log"
)

func InsertNote(db *sql.DB, name string, note string) {
	insertNoteSQL := `INSERT INTO notes(name, note) VALUES (?, ?)`
	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	_, err = statement.Exec(name, note)
	if err != nil {
		log.Println("Error inserting notes")
		return
	}
}
