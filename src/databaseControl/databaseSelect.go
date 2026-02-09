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
func SelectFromDBrow(db *sql.DB) []Test {
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

// selecting from db by name.  //завтра зробить!!!
func SelectFromDBbyName() {
	return
}
