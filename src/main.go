package main

import (
	"database/sql"
	"log"
	"net/http"
	databasecontrol "notes-api/databaseControl"

	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

type Notes struct {
	Name string
	Note string
}

func main() {
	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal("Error while openning DB", err)
	}
	defer db.Close()

	databasecontrol.CreateTable(db)

	router := gin.Default()

	router.POST("/postNote", func(c *gin.Context) {
		var notes Notes

		if err := c.ShouldBindJSON(&notes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad JSON request",
			})
			return
		}
		databasecontrol.InsertNote(db, notes.Name, notes.Note)

		c.JSON(200, gin.H{
			"result: ": "All good!",
		})
	})

	router.GET("/getNoteByName", func(c *gin.Context) {
		//if err := c.ShouldBind(&request); err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"error": "bad JSON request",
		//	})
		//	return
		//}
		//databasecontrol.SelectFromDBrow(db)
		c.JSON(http.StatusOK, gin.H{
			"status": "All good!",
		})
	})

	router.GET("/getAllNotes", func(c *gin.Context) {
		notes := databasecontrol.SelectFromDBrow(db)
		c.JSON(200, notes)
	})

	router.Run() //running on localhost:8080
}
