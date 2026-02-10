package main

import (
	"database/sql"
	"log"
	"net/http"
	databasecontrol "notes-api/databaseControl"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.POST("/postNote", func(c *gin.Context) {
		var notes Notes

		if err := c.ShouldBindJSON(&notes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "bad JSON request",
			})
			return
		}
		err := databasecontrol.InsertNote(db, notes.Name, notes.Note)
		if err != nil {
			c.JSON(404, gin.H{
				"status": "error while inserting information in /postNote",
			})
		}
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/getNoteByName", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gin.H{
				"status": "error, query parameter is required",
			})
			return
		}
		note, err := databasecontrol.SelectFromDbByName(db, name)
		if err != nil {
			log.Println("Error: ", err)
			return
		}
		c.JSON(200, note)
	})

	router.GET("/getAllNotes", func(c *gin.Context) {
		notes, err := databasecontrol.SelectFromDBallRow(db)
		if err != nil {
			c.JSON(404, gin.H{
				"status": "error while get information from /getAllNotes",
			})
		}
		c.JSON(200, notes)
	})

	router.Run() //running on localhost:8080
}
