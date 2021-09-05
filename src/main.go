package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var db *sql.DB = DatabaseClient()

func ResponseWriter(w http.ResponseWriter, error bool, data string) {
	if error {
		w.WriteHeader(http.StatusBadRequest)
	}

	response, _ := json.Marshal(&JSONResponse{
		Error: error,
		Data:  data,
	})
	w.Write(response)
}

func Logger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body := json.NewDecoder(r.Body)
	var data RequestData
	err := body.Decode(&data)

	if err != nil {
		ResponseWriter(w, true, "Invalid request: "+err.Error())
		return
	}

	if data.Status == 200 {
		ResponseWriter(w, false, "ok")
		return
	}

	id := DatabaseWriter(data)

	ResponseWriter(w, false, fmt.Sprint(id))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Logger)

	// db := DatabaseClient()

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func DatabaseClient() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DatabaseWriter(data RequestData) int {
	// db := DatabaseClient()

	var id int
	query := `
	INSERT INTO logging (date, url, method, status, user_id, headers, body, comment)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id;
	`

	err := db.QueryRow(query, data.Date, data.Url, data.Method, data.Status,
		data.UserID, data.Headers, data.Body, data.Comment).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	return id
}
