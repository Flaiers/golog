package main

import (
	"database/sql"
	"fmt"
	"os"

	"go-logging/src/config"
	"go-logging/src/log"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var db *sql.DB = DatabaseClient()

func DatabaseClient() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))

	if err != nil {
		log.Error(err)
	}

	return db
}

func DatabaseWriter(data config.RequestData) {
	var id int
	query := `
	INSERT INTO logging (date, url, method, status, user_id, body, comment)
	VALUES ($1, $2, $3, $4, $5, NULLIF($6, ''), NULLIF($7, ''))
	RETURNING id;
	`

	if err := db.QueryRow(query, data.Date, data.Url, data.Method, data.Status,
		data.UserID, data.Body, data.Comment).Scan(&id); err != nil {
		log.Info(err)
	}

	fmt.Println(id)
}
