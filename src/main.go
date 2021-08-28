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

var _ = godotenv.Load()

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: os.Getenv("REDIS_PASSWORD"),
	DB:       0,
})

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

	key := data.Date
	value, _ := json.Marshal(data)
	record := RedisWriter(key, string(value))

	if record != nil {
		ResponseWriter(w, true, "Failed write to redis: "+record.Error())
		return
	}

	ResponseWriter(w, false, "ok")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", Logger)

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}

func RedisWriter(key string, value string) error {
	err := client.Set(key, value, 0).Err()
	return err
}
