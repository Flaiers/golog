package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"error": false}`))
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
