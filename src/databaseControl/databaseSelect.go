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
func SelectFromDBallRow(db *sql.DB) []Test {
	var tests []Test
	query := "SELECT id, name, note FROM notes"
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error while selecting from DataBase: ", err)
	}
	defer rows.Close()

	var id int
	var name string
	var note string

	for rows.Next() {
		rows.Scan(&id, &name, &note)

		test := Test{ID: id, Name: name, Note: note}

		tests = append(tests, test)
	}

	return tests
}

// selecting from db by name.
func SelectFromDbByName(db *sql.DB, name string) ([]Test, error) {
	query := "SELECT * FROM notes WHERE name = ?"
	rows, err := db.Query(query, name)
	if err != nil {
		log.Println("Error while selecting from solo from DataBase: ", err)
	}
	defer rows.Close()

	tests := make([]Test, 0)

	for rows.Next() {
		var t Test
		if err := rows.Scan(&t.ID, &t.Name, &t.Note); err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tests, nil
}
