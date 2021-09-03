package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

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

	values, _ := json.Marshal(data)

	ResponseWriter(w, false, string(values))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Logger)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func DatabaseClient() {
	db, err := sql.Open("postgres", os.Getenv("DB_DSN"))

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()
}
