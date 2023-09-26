package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type note struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

var notes = []note{
	{ID: "1", Title: "FirstEntry", Content: "First not", CreatedAt: "2006-01-02"},
	{ID: "2", Title: "SecondEntry", Content: "Something more", CreatedAt: "2012-01-02"},
}

func getNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, notes)
}

func postNotes(c *gin.Context) {
	var newNote note
	if err := c.BindJSON(&newNote); err != nil {
		log.Fatal("binding json error", err)
		return
	}

	notes = append(notes, newNote)
	c.IndentedJSON(http.StatusOK, newNote)
}

func main() {
	router := gin.Default()
	router.GET("/notes", getNotes)
	router.POST("/addNotes", postNotes)

	router.Run("localhost:8080")
}

// curl http://localhost:8080/addNotes\
//   	--include \
//   	--header "Content-Type: application/json" \
// 	--request "POST" \
// 	--data '{"id": "4", "title": "Hello", "content":"world","created_at": "2013-02-01"}'
