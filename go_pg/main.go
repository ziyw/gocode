package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Note struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

func main() {

	connStr := "user=ziyan password=postgres dbname=ziyan sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	rows, err := db.Query("SELECT * FROM note")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var note Note
		var createdAt string
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &createdAt); err != nil {
			log.Fatal("unpack error", err)
		}

		date, err := time.Parse(time.RFC3339, createdAt)
		if err != nil {
			log.Fatal("parse date error", err)
		}
		note.CreatedAt = date

		fmt.Printf("ID: %d\tTitle: %s\t\t\tContent: %s, Created At: %s\n", note.ID, note.Title, note.Content, note.CreatedAt.Local().Format("2006-01-02"))
	}
	if err := rows.Err(); err != nil {
		log.Fatal("rows close error", err)
	}

}
