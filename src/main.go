package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: os.Getenv("REDIS_PASSWORD"),
	DB:       0,
})

type JSONResponse struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

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

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	ResponseWriter(w, false, "OK")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
