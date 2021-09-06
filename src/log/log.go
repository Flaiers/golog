package log

import (
	"log"
	"os"
)

func Info(message error) {
	info, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer info.Close()

	log := log.New(info, "INFO\t", log.LstdFlags)
	log.Print(message)
}

func Error(message error) {
	error, err := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer error.Close()

	log := log.New(error, "ERROR\t", log.LstdFlags)
	log.Fatal(message)
}
