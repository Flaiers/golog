package main

import (
	"encoding/json"
	"net/http"

	"go-logging/src/config"
	"go-logging/src/log"

	"github.com/gorilla/mux"
)

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
		ResponseWriter(w, false, "Did nothing")
		return
	}

	go DatabaseWriter(data)

	ResponseWriter(w, false, "ok")
}

func Counter(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Counter).Methods("GET")
	router.HandleFunc("/", Logger).Methods("POST")

	if err := db.Ping(); err != nil {
		log.Error(err)
	}
	defer db.Close()

	log.Error(http.ListenAndServe("0.0.0.0:8080", router))
}
