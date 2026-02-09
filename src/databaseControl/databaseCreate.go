package databasecontrol

import (
	"database/sql"
	"log"
)

func CreateTable(db *sql.DB) {
	createNotesTableSQL := `CREATE TABLE IF NOT EXISTS notes (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"name" TEXT NOT NULL UNIQUE,
		"note" TEXT NOT NULL UNIQUE
	);`

	log.Println("Creating users table...")
	statement, err := db.Prepare(createNotesTableSQL)
	if err != nil {
		log.Fatal("Error while creating table: ", err)
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		log.Println("Error:", err)
		return
	}
}
