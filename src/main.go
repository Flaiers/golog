package main

import (
	"encoding/json"
	"log"
	"net/http"

	"go-logging/src/config"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func ResponseWriter(w http.ResponseWriter, error bool, data string) {
	if error {
		w.WriteHeader(http.StatusBadRequest)
	}

	response, _ := json.Marshal(config.JSONResponse{
		Error: error,
		Data:  data,
	})
	w.Write(response)
}

func Logger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	body := json.NewDecoder(r.Body)
	var data config.RequestData

	if err := body.Decode(&data); err != nil {
		ResponseWriter(w, true, "Invalid request: "+err.Error())
		return
	}

	if data.Status == 200 {
		ResponseWriter(w, false, "ok")
		return
	}

	go DatabaseWriter(data)

	ResponseWriter(w, false, "ok")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Logger).Methods("POST")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
